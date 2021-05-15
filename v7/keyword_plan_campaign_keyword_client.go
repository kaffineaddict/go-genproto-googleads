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

var newKeywordPlanCampaignKeywordClientHook clientHook

// KeywordPlanCampaignKeywordCallOptions contains the retry settings for each method of KeywordPlanCampaignKeywordClient.
type KeywordPlanCampaignKeywordCallOptions struct {
	GetKeywordPlanCampaignKeyword     []gax.CallOption
	MutateKeywordPlanCampaignKeywords []gax.CallOption
}

func defaultKeywordPlanCampaignKeywordClientOptions() []option.ClientOption {
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

func defaultKeywordPlanCampaignKeywordCallOptions() *KeywordPlanCampaignKeywordCallOptions {
	return &KeywordPlanCampaignKeywordCallOptions{
		GetKeywordPlanCampaignKeyword: []gax.CallOption{
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
		MutateKeywordPlanCampaignKeywords: []gax.CallOption{
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

// KeywordPlanCampaignKeywordClient is a client for interacting with Google Ads API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type KeywordPlanCampaignKeywordClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	keywordPlanCampaignKeywordClient servicespb.KeywordPlanCampaignKeywordServiceClient

	// The call options for this service.
	CallOptions *KeywordPlanCampaignKeywordCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewKeywordPlanCampaignKeywordClient creates a new keyword plan campaign keyword service client.
//
// Service to manage Keyword Plan campaign keywords. KeywordPlanCampaign is
// required to add the campaign keywords. Only negative keywords are supported.
// A maximum of 1000 negative keywords are allowed per plan. This includes both
// campaign negative keywords and ad group negative keywords.
func NewKeywordPlanCampaignKeywordClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanCampaignKeywordClient, error) {
	clientOpts := defaultKeywordPlanCampaignKeywordClientOptions()

	if newKeywordPlanCampaignKeywordClientHook != nil {
		hookOpts, err := newKeywordPlanCampaignKeywordClientHook(ctx, clientHookParams{})
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
	c := &KeywordPlanCampaignKeywordClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultKeywordPlanCampaignKeywordCallOptions(),

		keywordPlanCampaignKeywordClient: servicespb.NewKeywordPlanCampaignKeywordServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *KeywordPlanCampaignKeywordClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanCampaignKeywordClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanCampaignKeywordClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// GetKeywordPlanCampaignKeyword returns the requested plan in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanCampaignKeywordClient) GetKeywordPlanCampaignKeyword(ctx context.Context, req *servicespb.GetKeywordPlanCampaignKeywordRequest, opts ...gax.CallOption) (*resourcespb.KeywordPlanCampaignKeyword, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetKeywordPlanCampaignKeyword[0:len(c.CallOptions.GetKeywordPlanCampaignKeyword):len(c.CallOptions.GetKeywordPlanCampaignKeyword)], opts...)
	var resp *resourcespb.KeywordPlanCampaignKeyword
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanCampaignKeywordClient.GetKeywordPlanCampaignKeyword(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MutateKeywordPlanCampaignKeywords creates, updates, or removes Keyword Plan campaign keywords. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanAdGroupKeywordError (at )
// KeywordPlanCampaignKeywordError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *KeywordPlanCampaignKeywordClient) MutateKeywordPlanCampaignKeywords(ctx context.Context, req *servicespb.MutateKeywordPlanCampaignKeywordsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanCampaignKeywordsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.MutateKeywordPlanCampaignKeywords[0:len(c.CallOptions.MutateKeywordPlanCampaignKeywords):len(c.CallOptions.MutateKeywordPlanCampaignKeywords)], opts...)
	var resp *servicespb.MutateKeywordPlanCampaignKeywordsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanCampaignKeywordClient.MutateKeywordPlanCampaignKeywords(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}