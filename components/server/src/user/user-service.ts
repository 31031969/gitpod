/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { inject, injectable } from "inversify";
import { Config } from "../config";
import { UserDB } from "@gitpod/gitpod-db/lib";
import { Authorizer } from "../authorization/authorizer";
import {
    AdditionalUserData,
    Identity,
    RoleOrPermission,
    TokenEntry,
    User,
    WorkspaceTimeoutDuration,
    WorkspaceTimeoutSetting,
} from "@gitpod/gitpod-protocol";
import { ApplicationError, ErrorCodes } from "@gitpod/gitpod-protocol/lib/messaging/error";
import { log } from "@gitpod/gitpod-protocol/lib/util/logging";
import { CreateUserParams } from "./user-authentication";
import { IAnalyticsWriter } from "@gitpod/gitpod-protocol/lib/analytics";
import { TransactionalContext } from "@gitpod/gitpod-db/lib/typeorm/transactional-db-impl";
import { RelationshipUpdater } from "../authorization/relationship-updater";

@injectable()
export class UserService {
    constructor(
        @inject(Config) private readonly config: Config,
        @inject(UserDB) private readonly userDb: UserDB,
        @inject(Authorizer) private readonly authorizer: Authorizer,
        @inject(IAnalyticsWriter) private readonly analytics: IAnalyticsWriter,
        @inject(RelationshipUpdater) private readonly relationshipUpdater: RelationshipUpdater,
    ) {}

    public async createUser(
        { organizationId, identity, token, userUpdate }: CreateUserParams,
        transactionCtx?: TransactionalContext,
    ): Promise<User> {
        log.debug("Creating new user.", { identity, "login-flow": true });
        return await this.userDb.transaction(transactionCtx, async (userDb) => {
            const newUser = await userDb.newUser();
            newUser.organizationId = organizationId;
            if (userUpdate) {
                userUpdate(newUser);
            }
            // HINT: we need to specify `deleted: false` here, so that any attempt to reuse the same
            // entry would converge to a valid state. The identities are identified by the external
            // `authId`, and if accounts are deleted, such entries are soft-deleted until the periodic
            // deleter will take care of them. Reuse of soft-deleted entries would lead to an invalid
            // state. This measure of prevention is considered in the period deleter as well.
            newUser.identities.push({ ...identity, deleted: false });
            this.handleNewUser(newUser);
            // new users should not see the migration message
            AdditionalUserData.set(newUser, { shouldSeeMigrationMessage: false });
            const result = await userDb.storeUser(newUser);
            await this.authorizer.addUser(result.id, organizationId);
            if (token) {
                await userDb.storeSingleToken(identity, token);
            }
            return result;
        });
    }

    private handleNewUser(newUser: User) {
        if (this.config.blockNewUsers.enabled) {
            const emailDomainInPasslist = (mail: string) =>
                this.config.blockNewUsers.passlist.some((e) => mail.endsWith(`@${e}`));
            const canPass = newUser.identities.some((i) => !!i.primaryEmail && emailDomainInPasslist(i.primaryEmail));

            // blocked = if user already blocked OR is not allowed to pass
            newUser.blocked = newUser.blocked || !canPass;
        }
    }

    async findUserById(userId: string, targetId: string): Promise<User> {
        const result = await this.userDb.findUserById(targetId);
        if (userId !== targetId) {
            await this.authorizer.checkPermissionOnUser(userId, "read_info", targetId, result);
        }
        if (!result) {
            throw new ApplicationError(ErrorCodes.NOT_FOUND, "not found");
        }

        try {
            return await this.relationshipUpdater.migrate(result);
        } catch (error) {
            log.error({ userId: targetId }, "Failed to migrate user", error);
            return result;
        }
    }

    async findTokensForIdentity(userId: string, identity: Identity): Promise<TokenEntry[]> {
        const result = await this.userDb.findTokensForIdentity(identity);
        return result;
    }

    async updateUser(userId: string, update: Partial<User> & { id: string }): Promise<User> {
        const targetUser = await this.findUserById(userId, update.id);
        await this.authorizer.checkPermissionOnUser(userId, "write_info", update.id, targetUser);

        //hang on to user profile before it's overwritten for analytics below
        const oldProfile = User.getProfile(targetUser);
        const allowedFields: (keyof User)[] = ["fullName", "additionalData"];
        for (const p of allowedFields) {
            if (p in update) {
                (targetUser[p] as any) = update[p];
            }
        }

        await this.userDb.updateUserPartial(targetUser);

        //track event and user profile if profile of partialUser changed
        const newProfile = User.getProfile(targetUser);
        if (User.Profile.hasChanges(oldProfile, newProfile)) {
            this.analytics.track({
                userId: targetUser.id,
                event: "profile_changed",
                properties: { new: newProfile, old: oldProfile },
            });
            this.analytics.identify({
                userId: targetUser.id,
                traits: { email: newProfile.email, company: newProfile.company, name: newProfile.name },
            });
        }
        return targetUser;
    }

    async updateWorkspaceTimeoutSetting(
        userId: string,
        targetUserId: string,
        setting: Partial<WorkspaceTimeoutSetting>,
    ): Promise<void> {
        if (setting.workspaceTimeout) {
            try {
                WorkspaceTimeoutDuration.validate(setting.workspaceTimeout);
            } catch (err) {
                throw new ApplicationError(ErrorCodes.BAD_REQUEST, err.message);
            }
        }

        const target = await this.findUserById(userId, targetUserId);
        await this.authorizer.checkPermissionOnUser(userId, "write_info", targetUserId, target);

        AdditionalUserData.set(target, setting);
        await this.userDb.updateUserPartial(target);
    }

    async listUsers(
        userId: string,
        req: {
            //
            offset?: number;
            limit?: number;
            orderBy?: keyof User;
            orderDir?: "ASC" | "DESC";
            searchTerm?: string;
        },
    ): Promise<{ total: number; rows: User[] }> {
        try {
            const res = await this.userDb.findAllUsers(
                req.offset || 0,
                req.limit || 100,
                req.orderBy || "creationDate",
                req.orderDir || "DESC",
                req.searchTerm,
            );
            const result = { total: res.total, rows: [] as User[] };
            for (const user of res.rows) {
                if (await this.authorizer.hasPermissionOnUser(userId, "read_info", user.id, user)) {
                    result.rows.push(user);
                } else {
                    result.total--;
                }
            }
            return result;
        } catch (e) {
            throw new ApplicationError(ErrorCodes.INTERNAL_SERVER_ERROR, e.toString());
        }
    }

    async updateRoleOrPermission(
        userId: string,
        targetUserId: string,
        modifications: { role: RoleOrPermission; add?: boolean }[],
    ): Promise<void> {
        const target = await this.findUserById(userId, targetUserId);
        await this.authorizer.checkPermissionOnUser(userId, "make_admin", targetUserId, target);

        const rolesOrPermissions = new Set((target.rolesOrPermissions || []) as string[]);
        const adminBefore = rolesOrPermissions.has("admin");
        modifications.forEach((e) => {
            if (e.add) {
                rolesOrPermissions.add(e.role as string);
            } else {
                rolesOrPermissions.delete(e.role as string);
            }
        });
        target.rolesOrPermissions = Array.from(rolesOrPermissions.values()) as RoleOrPermission[];
        const adminAfter = new Set(target.rolesOrPermissions).has("admin");
        try {
            await this.userDb.transaction(async (userDb) => {
                await userDb.storeUser(target);
                if (adminBefore !== adminAfter) {
                    if (adminAfter) {
                        await this.authorizer.addInstallationAdminRole(target.id);
                    } else {
                        await this.authorizer.removeInstallationAdminRole(target.id);
                    }
                }
            });
        } catch (error) {
            if (adminBefore !== adminAfter) {
                if (adminAfter) {
                    await this.authorizer.removeInstallationAdminRole(target.id);
                } else {
                    await this.authorizer.addInstallationAdminRole(target.id);
                }
            }
            throw error;
        }
    }

    async resetFgaVersion(subjectId: string, user: User) {
        await this.authorizer.checkPermissionOnUser(subjectId, "write_info", user.id, user);

        await this.userDb.updateUserPartial({ id: user.id, fgaRelationshipsVersion: undefined });
    }
}
