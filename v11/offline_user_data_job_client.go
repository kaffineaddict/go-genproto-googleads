// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v11/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v11/services"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newOfflineUserDataJobClientHook clientHook

// OfflineUserDataJobCallOptions contains the retry settings for each method of OfflineUserDataJobClient.
type OfflineUserDataJobCallOptions struct {
	CreateOfflineUserDataJob        []gax.CallOption
	AddOfflineUserDataJobOperations []gax.CallOption
	RunOfflineUserDataJob           []gax.CallOption
}

func defaultOfflineUserDataJobGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultOfflineUserDataJobCallOptions() *OfflineUserDataJobCallOptions {
	return &OfflineUserDataJobCallOptions{
		CreateOfflineUserDataJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AddOfflineUserDataJobOperations: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RunOfflineUserDataJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalOfflineUserDataJobClient is an interface that defines the methods available from Google Ads API.
type internalOfflineUserDataJobClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateOfflineUserDataJob(context.Context, *servicespb.CreateOfflineUserDataJobRequest, ...gax.CallOption) (*servicespb.CreateOfflineUserDataJobResponse, error)
	AddOfflineUserDataJobOperations(context.Context, *servicespb.AddOfflineUserDataJobOperationsRequest, ...gax.CallOption) (*servicespb.AddOfflineUserDataJobOperationsResponse, error)
	RunOfflineUserDataJob(context.Context, *servicespb.RunOfflineUserDataJobRequest, ...gax.CallOption) (*RunOfflineUserDataJobOperation, error)
	RunOfflineUserDataJobOperation(name string) *RunOfflineUserDataJobOperation
}

// OfflineUserDataJobClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage offline user data jobs.
type OfflineUserDataJobClient struct {
	// The internal transport-dependent client.
	internalClient internalOfflineUserDataJobClient

	// The call options for this service.
	CallOptions *OfflineUserDataJobCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *OfflineUserDataJobClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *OfflineUserDataJobClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *OfflineUserDataJobClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateOfflineUserDataJob creates an offline user data job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// NotAllowlistedError (at )
// OfflineUserDataJobError (at )
// QuotaError (at )
// RequestError (at )
func (c *OfflineUserDataJobClient) CreateOfflineUserDataJob(ctx context.Context, req *servicespb.CreateOfflineUserDataJobRequest, opts ...gax.CallOption) (*servicespb.CreateOfflineUserDataJobResponse, error) {
	return c.internalClient.CreateOfflineUserDataJob(ctx, req, opts...)
}

// AddOfflineUserDataJobOperations adds operations to the offline user data job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// OfflineUserDataJobError (at )
// QuotaError (at )
// RequestError (at )
func (c *OfflineUserDataJobClient) AddOfflineUserDataJobOperations(ctx context.Context, req *servicespb.AddOfflineUserDataJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddOfflineUserDataJobOperationsResponse, error) {
	return c.internalClient.AddOfflineUserDataJobOperations(ctx, req, opts...)
}

// RunOfflineUserDataJob runs the offline user data job.
//
// When finished, the long running operation will contain the processing
// result or failure information, if any.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// HeaderError (at )
// InternalError (at )
// OfflineUserDataJobError (at )
// QuotaError (at )
// RequestError (at )
func (c *OfflineUserDataJobClient) RunOfflineUserDataJob(ctx context.Context, req *servicespb.RunOfflineUserDataJobRequest, opts ...gax.CallOption) (*RunOfflineUserDataJobOperation, error) {
	return c.internalClient.RunOfflineUserDataJob(ctx, req, opts...)
}

// RunOfflineUserDataJobOperation returns a new RunOfflineUserDataJobOperation from a given name.
// The name must be that of a previously created RunOfflineUserDataJobOperation, possibly from a different process.
func (c *OfflineUserDataJobClient) RunOfflineUserDataJobOperation(name string) *RunOfflineUserDataJobOperation {
	return c.internalClient.RunOfflineUserDataJobOperation(name)
}

// offlineUserDataJobGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type offlineUserDataJobGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing OfflineUserDataJobClient
	CallOptions **OfflineUserDataJobCallOptions

	// The gRPC API client.
	offlineUserDataJobClient servicespb.OfflineUserDataJobServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewOfflineUserDataJobClient creates a new offline user data job service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage offline user data jobs.
func NewOfflineUserDataJobClient(ctx context.Context, opts ...option.ClientOption) (*OfflineUserDataJobClient, error) {
	clientOpts := defaultOfflineUserDataJobGRPCClientOptions()
	if newOfflineUserDataJobClientHook != nil {
		hookOpts, err := newOfflineUserDataJobClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := OfflineUserDataJobClient{CallOptions: defaultOfflineUserDataJobCallOptions()}

	c := &offlineUserDataJobGRPCClient{
		connPool:                 connPool,
		disableDeadlines:         disableDeadlines,
		offlineUserDataJobClient: servicespb.NewOfflineUserDataJobServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *offlineUserDataJobGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *offlineUserDataJobGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *offlineUserDataJobGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *offlineUserDataJobGRPCClient) CreateOfflineUserDataJob(ctx context.Context, req *servicespb.CreateOfflineUserDataJobRequest, opts ...gax.CallOption) (*servicespb.CreateOfflineUserDataJobResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateOfflineUserDataJob[0:len((*c.CallOptions).CreateOfflineUserDataJob):len((*c.CallOptions).CreateOfflineUserDataJob)], opts...)
	var resp *servicespb.CreateOfflineUserDataJobResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.offlineUserDataJobClient.CreateOfflineUserDataJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *offlineUserDataJobGRPCClient) AddOfflineUserDataJobOperations(ctx context.Context, req *servicespb.AddOfflineUserDataJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddOfflineUserDataJobOperationsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AddOfflineUserDataJobOperations[0:len((*c.CallOptions).AddOfflineUserDataJobOperations):len((*c.CallOptions).AddOfflineUserDataJobOperations)], opts...)
	var resp *servicespb.AddOfflineUserDataJobOperationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.offlineUserDataJobClient.AddOfflineUserDataJobOperations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *offlineUserDataJobGRPCClient) RunOfflineUserDataJob(ctx context.Context, req *servicespb.RunOfflineUserDataJobRequest, opts ...gax.CallOption) (*RunOfflineUserDataJobOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunOfflineUserDataJob[0:len((*c.CallOptions).RunOfflineUserDataJob):len((*c.CallOptions).RunOfflineUserDataJob)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.offlineUserDataJobClient.RunOfflineUserDataJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunOfflineUserDataJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// RunOfflineUserDataJobOperation manages a long-running operation from RunOfflineUserDataJob.
type RunOfflineUserDataJobOperation struct {
	lro *longrunning.Operation
}

// RunOfflineUserDataJobOperation returns a new RunOfflineUserDataJobOperation from a given name.
// The name must be that of a previously created RunOfflineUserDataJobOperation, possibly from a different process.
func (c *offlineUserDataJobGRPCClient) RunOfflineUserDataJobOperation(name string) *RunOfflineUserDataJobOperation {
	return &RunOfflineUserDataJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunOfflineUserDataJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunOfflineUserDataJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunOfflineUserDataJobOperation) Metadata() (*resourcespb.OfflineUserDataJobMetadata, error) {
	var meta resourcespb.OfflineUserDataJobMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunOfflineUserDataJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunOfflineUserDataJobOperation) Name() string {
	return op.lro.Name()
}
