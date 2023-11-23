// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/cluster.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ClusterServiceName is the fully-qualified name of the ClusterService service.
	ClusterServiceName = "gitpod.v1.ClusterService"
)

// ClusterServiceClient is a client for the gitpod.v1.ClusterService service.
type ClusterServiceClient interface {
	// RegisterCluster registers a new cluster with the server. Registrations are very
	// short-lived and must be renewed every 30 seconds. Clusters can be registered for
	// an entire organisation or a single user.
	RegisterCluster(context.Context, *connect_go.Request[v1.RegisterClusterRequest]) (*connect_go.Response[v1.RegisterClusterResponse], error)
	// RenewClusterRegistration renews a cluster's registration. This must be called every 30 seconds
	// to keep the cluster registered.
	RenewClusterRegistration(context.Context, *connect_go.Request[v1.RenewClusterRegistrationRequest]) (*connect_go.Response[v1.RenewClusterRegistrationResponse], error)
	// GetClusterWorkspaces returns the workspaces running on a cluster.
	GetClusterWorkspaces(context.Context, *connect_go.Request[v1.GetClusterWorkspacesRequest]) (*connect_go.Response[v1.GetClusterWorkspacesResponse], error)
	// UpdateClusterWorkspaceStatus updates the status of a workspace running on a cluster.
	UpdateClusterWorkspaceStatus(context.Context, *connect_go.Request[v1.UpdateClusterWorkspaceStatusRequest]) (*connect_go.Response[v1.UpdateClusterWorkspaceStatusResponse], error)
}

// NewClusterServiceClient constructs a client for the gitpod.v1.ClusterService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ClusterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clusterServiceClient{
		registerCluster: connect_go.NewClient[v1.RegisterClusterRequest, v1.RegisterClusterResponse](
			httpClient,
			baseURL+"/gitpod.v1.ClusterService/RegisterCluster",
			opts...,
		),
		renewClusterRegistration: connect_go.NewClient[v1.RenewClusterRegistrationRequest, v1.RenewClusterRegistrationResponse](
			httpClient,
			baseURL+"/gitpod.v1.ClusterService/RenewClusterRegistration",
			opts...,
		),
		getClusterWorkspaces: connect_go.NewClient[v1.GetClusterWorkspacesRequest, v1.GetClusterWorkspacesResponse](
			httpClient,
			baseURL+"/gitpod.v1.ClusterService/GetClusterWorkspaces",
			opts...,
		),
		updateClusterWorkspaceStatus: connect_go.NewClient[v1.UpdateClusterWorkspaceStatusRequest, v1.UpdateClusterWorkspaceStatusResponse](
			httpClient,
			baseURL+"/gitpod.v1.ClusterService/UpdateClusterWorkspaceStatus",
			opts...,
		),
	}
}

// clusterServiceClient implements ClusterServiceClient.
type clusterServiceClient struct {
	registerCluster              *connect_go.Client[v1.RegisterClusterRequest, v1.RegisterClusterResponse]
	renewClusterRegistration     *connect_go.Client[v1.RenewClusterRegistrationRequest, v1.RenewClusterRegistrationResponse]
	getClusterWorkspaces         *connect_go.Client[v1.GetClusterWorkspacesRequest, v1.GetClusterWorkspacesResponse]
	updateClusterWorkspaceStatus *connect_go.Client[v1.UpdateClusterWorkspaceStatusRequest, v1.UpdateClusterWorkspaceStatusResponse]
}

// RegisterCluster calls gitpod.v1.ClusterService.RegisterCluster.
func (c *clusterServiceClient) RegisterCluster(ctx context.Context, req *connect_go.Request[v1.RegisterClusterRequest]) (*connect_go.Response[v1.RegisterClusterResponse], error) {
	return c.registerCluster.CallUnary(ctx, req)
}

// RenewClusterRegistration calls gitpod.v1.ClusterService.RenewClusterRegistration.
func (c *clusterServiceClient) RenewClusterRegistration(ctx context.Context, req *connect_go.Request[v1.RenewClusterRegistrationRequest]) (*connect_go.Response[v1.RenewClusterRegistrationResponse], error) {
	return c.renewClusterRegistration.CallUnary(ctx, req)
}

// GetClusterWorkspaces calls gitpod.v1.ClusterService.GetClusterWorkspaces.
func (c *clusterServiceClient) GetClusterWorkspaces(ctx context.Context, req *connect_go.Request[v1.GetClusterWorkspacesRequest]) (*connect_go.Response[v1.GetClusterWorkspacesResponse], error) {
	return c.getClusterWorkspaces.CallUnary(ctx, req)
}

// UpdateClusterWorkspaceStatus calls gitpod.v1.ClusterService.UpdateClusterWorkspaceStatus.
func (c *clusterServiceClient) UpdateClusterWorkspaceStatus(ctx context.Context, req *connect_go.Request[v1.UpdateClusterWorkspaceStatusRequest]) (*connect_go.Response[v1.UpdateClusterWorkspaceStatusResponse], error) {
	return c.updateClusterWorkspaceStatus.CallUnary(ctx, req)
}

// ClusterServiceHandler is an implementation of the gitpod.v1.ClusterService service.
type ClusterServiceHandler interface {
	// RegisterCluster registers a new cluster with the server. Registrations are very
	// short-lived and must be renewed every 30 seconds. Clusters can be registered for
	// an entire organisation or a single user.
	RegisterCluster(context.Context, *connect_go.Request[v1.RegisterClusterRequest]) (*connect_go.Response[v1.RegisterClusterResponse], error)
	// RenewClusterRegistration renews a cluster's registration. This must be called every 30 seconds
	// to keep the cluster registered.
	RenewClusterRegistration(context.Context, *connect_go.Request[v1.RenewClusterRegistrationRequest]) (*connect_go.Response[v1.RenewClusterRegistrationResponse], error)
	// GetClusterWorkspaces returns the workspaces running on a cluster.
	GetClusterWorkspaces(context.Context, *connect_go.Request[v1.GetClusterWorkspacesRequest]) (*connect_go.Response[v1.GetClusterWorkspacesResponse], error)
	// UpdateClusterWorkspaceStatus updates the status of a workspace running on a cluster.
	UpdateClusterWorkspaceStatus(context.Context, *connect_go.Request[v1.UpdateClusterWorkspaceStatusRequest]) (*connect_go.Response[v1.UpdateClusterWorkspaceStatusResponse], error)
}

// NewClusterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterServiceHandler(svc ClusterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.ClusterService/RegisterCluster", connect_go.NewUnaryHandler(
		"/gitpod.v1.ClusterService/RegisterCluster",
		svc.RegisterCluster,
		opts...,
	))
	mux.Handle("/gitpod.v1.ClusterService/RenewClusterRegistration", connect_go.NewUnaryHandler(
		"/gitpod.v1.ClusterService/RenewClusterRegistration",
		svc.RenewClusterRegistration,
		opts...,
	))
	mux.Handle("/gitpod.v1.ClusterService/GetClusterWorkspaces", connect_go.NewUnaryHandler(
		"/gitpod.v1.ClusterService/GetClusterWorkspaces",
		svc.GetClusterWorkspaces,
		opts...,
	))
	mux.Handle("/gitpod.v1.ClusterService/UpdateClusterWorkspaceStatus", connect_go.NewUnaryHandler(
		"/gitpod.v1.ClusterService/UpdateClusterWorkspaceStatus",
		svc.UpdateClusterWorkspaceStatus,
		opts...,
	))
	return "/gitpod.v1.ClusterService/", mux
}

// UnimplementedClusterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterServiceHandler struct{}

func (UnimplementedClusterServiceHandler) RegisterCluster(context.Context, *connect_go.Request[v1.RegisterClusterRequest]) (*connect_go.Response[v1.RegisterClusterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.ClusterService.RegisterCluster is not implemented"))
}

func (UnimplementedClusterServiceHandler) RenewClusterRegistration(context.Context, *connect_go.Request[v1.RenewClusterRegistrationRequest]) (*connect_go.Response[v1.RenewClusterRegistrationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.ClusterService.RenewClusterRegistration is not implemented"))
}

func (UnimplementedClusterServiceHandler) GetClusterWorkspaces(context.Context, *connect_go.Request[v1.GetClusterWorkspacesRequest]) (*connect_go.Response[v1.GetClusterWorkspacesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.ClusterService.GetClusterWorkspaces is not implemented"))
}

func (UnimplementedClusterServiceHandler) UpdateClusterWorkspaceStatus(context.Context, *connect_go.Request[v1.UpdateClusterWorkspaceStatusRequest]) (*connect_go.Response[v1.UpdateClusterWorkspaceStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.ClusterService.UpdateClusterWorkspaceStatus is not implemented"))
}
