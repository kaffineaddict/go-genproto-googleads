// Copyright 2021 Google LLC
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

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "google.golang.org/genproto/googleapis/ads/googleads/v7/resources"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v7/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAdGroupLabelClientHook clientHook

// AdGroupLabelCallOptions contains the retry settings for each method of AdGroupLabelClient.
type AdGroupLabelCallOptions struct {
	GetAdGroupLabel     []gax.CallOption
	MutateAdGroupLabels []gax.CallOption
}

func defaultAdGroupLabelClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAdGroupLabelCallOptions() *AdGroupLabelCallOptions {
	return &AdGroupLabelCallOptions{
		GetAdGroupLabel: []gax.CallOption{
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
		MutateAdGroupLabels: []gax.CallOption{
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

// AdGroupLabelClient is a client for interacting with Google Ads API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type AdGroupLabelClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	adGroupLabelClient servicespb.AdGroupLabelServiceClient

	// The call options for this service.
	CallOptions *AdGroupLabelCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupLabelClient creates a new ad group label service client.
//
// Service to manage labels on ad groups.
func NewAdGroupLabelClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupLabelClient, error) {
	clientOpts := defaultAdGroupLabelClientOptions()

	if newAdGroupLabelClientHook != nil {
		hookOpts, err := newAdGroupLabelClientHook(ctx, clientHookParams{})
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
	c := &AdGroupLabelClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultAdGroupLabelCallOptions(),

		adGroupLabelClient: servicespb.NewAdGroupLabelServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupLabelClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupLabelClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupLabelClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// GetAdGroupLabel returns the requested ad group label in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupLabelClient) GetAdGroupLabel(ctx context.Context, req *servicespb.GetAdGroupLabelRequest, opts ...gax.CallOption) (*resourcespb.AdGroupLabel, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetAdGroupLabel[0:len(c.CallOptions.GetAdGroupLabel):len(c.CallOptions.GetAdGroupLabel)], opts...)
	var resp *resourcespb.AdGroupLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupLabelClient.GetAdGroupLabel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MutateAdGroupLabels creates and removes ad group labels.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// LabelError (at )
// MutateError (at )
// NewResourceCreationError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupLabelClient) MutateAdGroupLabels(ctx context.Context, req *servicespb.MutateAdGroupLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupLabelsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.MutateAdGroupLabels[0:len(c.CallOptions.MutateAdGroupLabels):len(c.CallOptions.MutateAdGroupLabels)], opts...)
	var resp *servicespb.MutateAdGroupLabelsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupLabelClient.MutateAdGroupLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}