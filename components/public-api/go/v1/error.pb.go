// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/v1/error.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// details for PERMISSION_DENIED status code
type PermissionDeniedDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reason:
	//
	//	*PermissionDeniedDetails_UserBlocked
	//	*PermissionDeniedDetails_NeedsVerification
	Reason isPermissionDeniedDetails_Reason `protobuf_oneof:"reason"`
}

func (x *PermissionDeniedDetails) Reset() {
	*x = PermissionDeniedDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionDeniedDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDeniedDetails) ProtoMessage() {}

func (x *PermissionDeniedDetails) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDeniedDetails.ProtoReflect.Descriptor instead.
func (*PermissionDeniedDetails) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{0}
}

func (m *PermissionDeniedDetails) GetReason() isPermissionDeniedDetails_Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (x *PermissionDeniedDetails) GetUserBlocked() *UserBlockedError {
	if x, ok := x.GetReason().(*PermissionDeniedDetails_UserBlocked); ok {
		return x.UserBlocked
	}
	return nil
}

func (x *PermissionDeniedDetails) GetNeedsVerification() *NeedsVerificationError {
	if x, ok := x.GetReason().(*PermissionDeniedDetails_NeedsVerification); ok {
		return x.NeedsVerification
	}
	return nil
}

type isPermissionDeniedDetails_Reason interface {
	isPermissionDeniedDetails_Reason()
}

type PermissionDeniedDetails_UserBlocked struct {
	UserBlocked *UserBlockedError `protobuf:"bytes,1,opt,name=user_blocked,json=userBlocked,proto3,oneof"`
}

type PermissionDeniedDetails_NeedsVerification struct {
	NeedsVerification *NeedsVerificationError `protobuf:"bytes,2,opt,name=needs_verification,json=needsVerification,proto3,oneof"`
}

func (*PermissionDeniedDetails_UserBlocked) isPermissionDeniedDetails_Reason() {}

func (*PermissionDeniedDetails_NeedsVerification) isPermissionDeniedDetails_Reason() {}

type UserBlockedError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserBlockedError) Reset() {
	*x = UserBlockedError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBlockedError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBlockedError) ProtoMessage() {}

func (x *UserBlockedError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBlockedError.ProtoReflect.Descriptor instead.
func (*UserBlockedError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{1}
}

type NeedsVerificationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NeedsVerificationError) Reset() {
	*x = NeedsVerificationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeedsVerificationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeedsVerificationError) ProtoMessage() {}

func (x *NeedsVerificationError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeedsVerificationError.ProtoReflect.Descriptor instead.
func (*NeedsVerificationError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{2}
}

// details for FAILED_PRECONDITION status code
type FailedPreconditionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reason:
	//
	//	*FailedPreconditionDetails_PaymentSpendingLimitReached
	//	*FailedPreconditionDetails_InvalidCostCenter
	//	*FailedPreconditionDetails_TooManyRunningWorkspaces
	//	*FailedPreconditionDetails_InvalidGitpodYml
	//	*FailedPreconditionDetails_RepositoryNotFound
	//	*FailedPreconditionDetails_RepositoryUnauthorized
	//	*FailedPreconditionDetails_ImageBuildLogsNotYetAvailable
	Reason isFailedPreconditionDetails_Reason `protobuf_oneof:"reason"`
}

func (x *FailedPreconditionDetails) Reset() {
	*x = FailedPreconditionDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedPreconditionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedPreconditionDetails) ProtoMessage() {}

func (x *FailedPreconditionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedPreconditionDetails.ProtoReflect.Descriptor instead.
func (*FailedPreconditionDetails) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{3}
}

func (m *FailedPreconditionDetails) GetReason() isFailedPreconditionDetails_Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (x *FailedPreconditionDetails) GetPaymentSpendingLimitReached() *PaymentSpendingLimitReachedError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_PaymentSpendingLimitReached); ok {
		return x.PaymentSpendingLimitReached
	}
	return nil
}

func (x *FailedPreconditionDetails) GetInvalidCostCenter() *InvalidCostCenterError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_InvalidCostCenter); ok {
		return x.InvalidCostCenter
	}
	return nil
}

func (x *FailedPreconditionDetails) GetTooManyRunningWorkspaces() *TooManyRunningWorkspacesError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_TooManyRunningWorkspaces); ok {
		return x.TooManyRunningWorkspaces
	}
	return nil
}

func (x *FailedPreconditionDetails) GetInvalidGitpodYml() *InvalidGitpodYMLError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_InvalidGitpodYml); ok {
		return x.InvalidGitpodYml
	}
	return nil
}

func (x *FailedPreconditionDetails) GetRepositoryNotFound() *RepositoryNotFoundError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_RepositoryNotFound); ok {
		return x.RepositoryNotFound
	}
	return nil
}

func (x *FailedPreconditionDetails) GetRepositoryUnauthorized() *RepositoryUnauthorizedError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_RepositoryUnauthorized); ok {
		return x.RepositoryUnauthorized
	}
	return nil
}

func (x *FailedPreconditionDetails) GetImageBuildLogsNotYetAvailable() *ImageBuildLogsNotYetAvailableError {
	if x, ok := x.GetReason().(*FailedPreconditionDetails_ImageBuildLogsNotYetAvailable); ok {
		return x.ImageBuildLogsNotYetAvailable
	}
	return nil
}

type isFailedPreconditionDetails_Reason interface {
	isFailedPreconditionDetails_Reason()
}

type FailedPreconditionDetails_PaymentSpendingLimitReached struct {
	PaymentSpendingLimitReached *PaymentSpendingLimitReachedError `protobuf:"bytes,1,opt,name=payment_spending_limit_reached,json=paymentSpendingLimitReached,proto3,oneof"`
}

type FailedPreconditionDetails_InvalidCostCenter struct {
	InvalidCostCenter *InvalidCostCenterError `protobuf:"bytes,2,opt,name=invalid_cost_center,json=invalidCostCenter,proto3,oneof"`
}

type FailedPreconditionDetails_TooManyRunningWorkspaces struct {
	TooManyRunningWorkspaces *TooManyRunningWorkspacesError `protobuf:"bytes,3,opt,name=too_many_running_workspaces,json=tooManyRunningWorkspaces,proto3,oneof"`
}

type FailedPreconditionDetails_InvalidGitpodYml struct {
	InvalidGitpodYml *InvalidGitpodYMLError `protobuf:"bytes,4,opt,name=invalid_gitpod_yml,json=invalidGitpodYml,proto3,oneof"`
}

type FailedPreconditionDetails_RepositoryNotFound struct {
	RepositoryNotFound *RepositoryNotFoundError `protobuf:"bytes,5,opt,name=repository_not_found,json=repositoryNotFound,proto3,oneof"`
}

type FailedPreconditionDetails_RepositoryUnauthorized struct {
	RepositoryUnauthorized *RepositoryUnauthorizedError `protobuf:"bytes,6,opt,name=repository_unauthorized,json=repositoryUnauthorized,proto3,oneof"`
}

type FailedPreconditionDetails_ImageBuildLogsNotYetAvailable struct {
	ImageBuildLogsNotYetAvailable *ImageBuildLogsNotYetAvailableError `protobuf:"bytes,7,opt,name=image_build_logs_not_yet_available,json=imageBuildLogsNotYetAvailable,proto3,oneof"`
}

func (*FailedPreconditionDetails_PaymentSpendingLimitReached) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_InvalidCostCenter) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_TooManyRunningWorkspaces) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_InvalidGitpodYml) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_RepositoryNotFound) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_RepositoryUnauthorized) isFailedPreconditionDetails_Reason() {}

func (*FailedPreconditionDetails_ImageBuildLogsNotYetAvailable) isFailedPreconditionDetails_Reason() {
}

type PaymentSpendingLimitReachedError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentSpendingLimitReachedError) Reset() {
	*x = PaymentSpendingLimitReachedError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentSpendingLimitReachedError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentSpendingLimitReachedError) ProtoMessage() {}

func (x *PaymentSpendingLimitReachedError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentSpendingLimitReachedError.ProtoReflect.Descriptor instead.
func (*PaymentSpendingLimitReachedError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{4}
}

type InvalidCostCenterError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributionId string `protobuf:"bytes,1,opt,name=attribution_id,json=attributionId,proto3" json:"attribution_id,omitempty"`
}

func (x *InvalidCostCenterError) Reset() {
	*x = InvalidCostCenterError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidCostCenterError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidCostCenterError) ProtoMessage() {}

func (x *InvalidCostCenterError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidCostCenterError.ProtoReflect.Descriptor instead.
func (*InvalidCostCenterError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{5}
}

func (x *InvalidCostCenterError) GetAttributionId() string {
	if x != nil {
		return x.AttributionId
	}
	return ""
}

type TooManyRunningWorkspacesError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TooManyRunningWorkspacesError) Reset() {
	*x = TooManyRunningWorkspacesError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TooManyRunningWorkspacesError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TooManyRunningWorkspacesError) ProtoMessage() {}

func (x *TooManyRunningWorkspacesError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TooManyRunningWorkspacesError.ProtoReflect.Descriptor instead.
func (*TooManyRunningWorkspacesError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{6}
}

type InvalidGitpodYMLError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Violations []string `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *InvalidGitpodYMLError) Reset() {
	*x = InvalidGitpodYMLError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidGitpodYMLError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidGitpodYMLError) ProtoMessage() {}

func (x *InvalidGitpodYMLError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidGitpodYMLError.ProtoReflect.Descriptor instead.
func (*InvalidGitpodYMLError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{7}
}

func (x *InvalidGitpodYMLError) GetViolations() []string {
	if x != nil {
		return x.Violations
	}
	return nil
}

type RepositoryNotFoundError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host        string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Owner       string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	UserIsOwner bool     `protobuf:"varint,3,opt,name=user_is_owner,json=userIsOwner,proto3" json:"user_is_owner,omitempty"`
	UserScopes  []string `protobuf:"bytes,4,rep,name=user_scopes,json=userScopes,proto3" json:"user_scopes,omitempty"`
	LastUpdate  string   `protobuf:"bytes,5,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *RepositoryNotFoundError) Reset() {
	*x = RepositoryNotFoundError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryNotFoundError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryNotFoundError) ProtoMessage() {}

func (x *RepositoryNotFoundError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryNotFoundError.ProtoReflect.Descriptor instead.
func (*RepositoryNotFoundError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{8}
}

func (x *RepositoryNotFoundError) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RepositoryNotFoundError) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RepositoryNotFoundError) GetUserIsOwner() bool {
	if x != nil {
		return x.UserIsOwner
	}
	return false
}

func (x *RepositoryNotFoundError) GetUserScopes() []string {
	if x != nil {
		return x.UserScopes
	}
	return nil
}

func (x *RepositoryNotFoundError) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type RepositoryUnauthorizedError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	RequiredScopes      []string `protobuf:"bytes,2,rep,name=required_scopes,json=requiredScopes,proto3" json:"required_scopes,omitempty"`
	ProviderType        string   `protobuf:"bytes,3,opt,name=provider_type,json=providerType,proto3" json:"provider_type,omitempty"`
	RepoName            string   `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	ProviderIsConnected bool     `protobuf:"varint,5,opt,name=provider_is_connected,json=providerIsConnected,proto3" json:"provider_is_connected,omitempty"`
}

func (x *RepositoryUnauthorizedError) Reset() {
	*x = RepositoryUnauthorizedError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryUnauthorizedError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryUnauthorizedError) ProtoMessage() {}

func (x *RepositoryUnauthorizedError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryUnauthorizedError.ProtoReflect.Descriptor instead.
func (*RepositoryUnauthorizedError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{9}
}

func (x *RepositoryUnauthorizedError) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RepositoryUnauthorizedError) GetRequiredScopes() []string {
	if x != nil {
		return x.RequiredScopes
	}
	return nil
}

func (x *RepositoryUnauthorizedError) GetProviderType() string {
	if x != nil {
		return x.ProviderType
	}
	return ""
}

func (x *RepositoryUnauthorizedError) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *RepositoryUnauthorizedError) GetProviderIsConnected() bool {
	if x != nil {
		return x.ProviderIsConnected
	}
	return false
}

type ImageBuildLogsNotYetAvailableError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImageBuildLogsNotYetAvailableError) Reset() {
	*x = ImageBuildLogsNotYetAvailableError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_error_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageBuildLogsNotYetAvailableError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageBuildLogsNotYetAvailableError) ProtoMessage() {}

func (x *ImageBuildLogsNotYetAvailableError) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_error_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageBuildLogsNotYetAvailableError.ProtoReflect.Descriptor instead.
func (*ImageBuildLogsNotYetAvailableError) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_error_proto_rawDescGZIP(), []int{10}
}

var File_gitpod_v1_error_proto protoreflect.FileDescriptor

var file_gitpod_v1_error_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x40, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x52, 0x0a, 0x12, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x73,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x11, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0x12, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x4e, 0x65, 0x65, 0x64, 0x73, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xe2,
	0x05, 0x0a, 0x19, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x72, 0x0a, 0x1e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x12, 0x53, 0x0a, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x1b, 0x74, 0x6f, 0x6f, 0x5f, 0x6d, 0x61, 0x6e,
	0x79, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x18, 0x74, 0x6f, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x52,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x50, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x5f, 0x79, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x47, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x59, 0x4d, 0x4c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x47, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x59,
	0x6d, 0x6c, 0x12, 0x56, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x61, 0x0a, 0x17, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69,
	0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x16, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x7a, 0x0a,
	0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x79, 0x65, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x6f, 0x74, 0x59, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x1d, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x4e, 0x6f, 0x74, 0x59, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x20, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x54, 0x6f, 0x6f, 0x4d,
	0x61, 0x6e, 0x79, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x37, 0x0a, 0x15, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x47, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x59, 0x4d, 0x4c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x73, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0xd0,
	0x01, 0x0a, 0x1b, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a,
	0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x22, 0x24, 0x0a, 0x22, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c,
	0x6f, 0x67, 0x73, 0x4e, 0x6f, 0x74, 0x59, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_v1_error_proto_rawDescOnce sync.Once
	file_gitpod_v1_error_proto_rawDescData = file_gitpod_v1_error_proto_rawDesc
)

func file_gitpod_v1_error_proto_rawDescGZIP() []byte {
	file_gitpod_v1_error_proto_rawDescOnce.Do(func() {
		file_gitpod_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_v1_error_proto_rawDescData)
	})
	return file_gitpod_v1_error_proto_rawDescData
}

var file_gitpod_v1_error_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_gitpod_v1_error_proto_goTypes = []interface{}{
	(*PermissionDeniedDetails)(nil),            // 0: gitpod.v1.PermissionDeniedDetails
	(*UserBlockedError)(nil),                   // 1: gitpod.v1.UserBlockedError
	(*NeedsVerificationError)(nil),             // 2: gitpod.v1.NeedsVerificationError
	(*FailedPreconditionDetails)(nil),          // 3: gitpod.v1.FailedPreconditionDetails
	(*PaymentSpendingLimitReachedError)(nil),   // 4: gitpod.v1.PaymentSpendingLimitReachedError
	(*InvalidCostCenterError)(nil),             // 5: gitpod.v1.InvalidCostCenterError
	(*TooManyRunningWorkspacesError)(nil),      // 6: gitpod.v1.TooManyRunningWorkspacesError
	(*InvalidGitpodYMLError)(nil),              // 7: gitpod.v1.InvalidGitpodYMLError
	(*RepositoryNotFoundError)(nil),            // 8: gitpod.v1.RepositoryNotFoundError
	(*RepositoryUnauthorizedError)(nil),        // 9: gitpod.v1.RepositoryUnauthorizedError
	(*ImageBuildLogsNotYetAvailableError)(nil), // 10: gitpod.v1.ImageBuildLogsNotYetAvailableError
}
var file_gitpod_v1_error_proto_depIdxs = []int32{
	1,  // 0: gitpod.v1.PermissionDeniedDetails.user_blocked:type_name -> gitpod.v1.UserBlockedError
	2,  // 1: gitpod.v1.PermissionDeniedDetails.needs_verification:type_name -> gitpod.v1.NeedsVerificationError
	4,  // 2: gitpod.v1.FailedPreconditionDetails.payment_spending_limit_reached:type_name -> gitpod.v1.PaymentSpendingLimitReachedError
	5,  // 3: gitpod.v1.FailedPreconditionDetails.invalid_cost_center:type_name -> gitpod.v1.InvalidCostCenterError
	6,  // 4: gitpod.v1.FailedPreconditionDetails.too_many_running_workspaces:type_name -> gitpod.v1.TooManyRunningWorkspacesError
	7,  // 5: gitpod.v1.FailedPreconditionDetails.invalid_gitpod_yml:type_name -> gitpod.v1.InvalidGitpodYMLError
	8,  // 6: gitpod.v1.FailedPreconditionDetails.repository_not_found:type_name -> gitpod.v1.RepositoryNotFoundError
	9,  // 7: gitpod.v1.FailedPreconditionDetails.repository_unauthorized:type_name -> gitpod.v1.RepositoryUnauthorizedError
	10, // 8: gitpod.v1.FailedPreconditionDetails.image_build_logs_not_yet_available:type_name -> gitpod.v1.ImageBuildLogsNotYetAvailableError
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_gitpod_v1_error_proto_init() }
func file_gitpod_v1_error_proto_init() {
	if File_gitpod_v1_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_v1_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionDeniedDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBlockedError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeedsVerificationError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedPreconditionDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentSpendingLimitReachedError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidCostCenterError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TooManyRunningWorkspacesError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidGitpodYMLError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryNotFoundError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryUnauthorizedError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_v1_error_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageBuildLogsNotYetAvailableError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_gitpod_v1_error_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PermissionDeniedDetails_UserBlocked)(nil),
		(*PermissionDeniedDetails_NeedsVerification)(nil),
	}
	file_gitpod_v1_error_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*FailedPreconditionDetails_PaymentSpendingLimitReached)(nil),
		(*FailedPreconditionDetails_InvalidCostCenter)(nil),
		(*FailedPreconditionDetails_TooManyRunningWorkspaces)(nil),
		(*FailedPreconditionDetails_InvalidGitpodYml)(nil),
		(*FailedPreconditionDetails_RepositoryNotFound)(nil),
		(*FailedPreconditionDetails_RepositoryUnauthorized)(nil),
		(*FailedPreconditionDetails_ImageBuildLogsNotYetAvailable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gitpod_v1_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gitpod_v1_error_proto_goTypes,
		DependencyIndexes: file_gitpod_v1_error_proto_depIdxs,
		MessageInfos:      file_gitpod_v1_error_proto_msgTypes,
	}.Build()
	File_gitpod_v1_error_proto = out.File
	file_gitpod_v1_error_proto_rawDesc = nil
	file_gitpod_v1_error_proto_goTypes = nil
	file_gitpod_v1_error_proto_depIdxs = nil
}
