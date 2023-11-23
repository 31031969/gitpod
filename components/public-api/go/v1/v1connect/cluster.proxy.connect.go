// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/v1"
)

var _ ClusterServiceHandler = (*ProxyClusterServiceHandler)(nil)

type ProxyClusterServiceHandler struct {
	Client v1.ClusterServiceClient
	UnimplementedClusterServiceHandler
}

func (s *ProxyClusterServiceHandler) RegisterCluster(ctx context.Context, req *connect_go.Request[v1.RegisterClusterRequest]) (*connect_go.Response[v1.RegisterClusterResponse], error) {
	resp, err := s.Client.RegisterCluster(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyClusterServiceHandler) RenewClusterRegistration(ctx context.Context, req *connect_go.Request[v1.RenewClusterRegistrationRequest]) (*connect_go.Response[v1.RenewClusterRegistrationResponse], error) {
	resp, err := s.Client.RenewClusterRegistration(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyClusterServiceHandler) GetClusterWorkspaces(ctx context.Context, req *connect_go.Request[v1.GetClusterWorkspacesRequest]) (*connect_go.Response[v1.GetClusterWorkspacesResponse], error) {
	resp, err := s.Client.GetClusterWorkspaces(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyClusterServiceHandler) UpdateClusterWorkspaceStatus(ctx context.Context, req *connect_go.Request[v1.UpdateClusterWorkspaceStatusRequest]) (*connect_go.Response[v1.UpdateClusterWorkspaceStatusResponse], error) {
	resp, err := s.Client.UpdateClusterWorkspaceStatus(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
