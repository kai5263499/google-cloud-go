// Copyright 2022 Google LLC
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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newInterconnectAttachmentsClientHook clientHook

// InterconnectAttachmentsCallOptions contains the retry settings for each method of InterconnectAttachmentsClient.
type InterconnectAttachmentsCallOptions struct {
	AggregatedList []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	Insert         []gax.CallOption
	List           []gax.CallOption
	Patch          []gax.CallOption
	SetLabels      []gax.CallOption
}

func defaultInterconnectAttachmentsRESTCallOptions() *InterconnectAttachmentsCallOptions {
	return &InterconnectAttachmentsCallOptions{
		AggregatedList: []gax.CallOption{},
		Delete:         []gax.CallOption{},
		Get:            []gax.CallOption{},
		Insert:         []gax.CallOption{},
		List:           []gax.CallOption{},
		Patch:          []gax.CallOption{},
		SetLabels:      []gax.CallOption{},
	}
}

// internalInterconnectAttachmentsClient is an interface that defines the methods available from Google Compute Engine API.
type internalInterconnectAttachmentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListInterconnectAttachmentsRequest, ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteInterconnectAttachmentRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.InterconnectAttachment, error)
	Insert(context.Context, *computepb.InsertInterconnectAttachmentRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListInterconnectAttachmentsRequest, ...gax.CallOption) *InterconnectAttachmentIterator
	Patch(context.Context, *computepb.PatchInterconnectAttachmentRequest, ...gax.CallOption) (*Operation, error)
	SetLabels(context.Context, *computepb.SetLabelsInterconnectAttachmentRequest, ...gax.CallOption) (*Operation, error)
}

// InterconnectAttachmentsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The InterconnectAttachments API.
type InterconnectAttachmentsClient struct {
	// The internal transport-dependent client.
	internalClient internalInterconnectAttachmentsClient

	// The call options for this service.
	CallOptions *InterconnectAttachmentsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InterconnectAttachmentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InterconnectAttachmentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InterconnectAttachmentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of interconnect attachments.
func (c *InterconnectAttachmentsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified interconnect attachment.
func (c *InterconnectAttachmentsClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified interconnect attachment.
func (c *InterconnectAttachmentsClient) Get(ctx context.Context, req *computepb.GetInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.InterconnectAttachment, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates an InterconnectAttachment in the specified project using the data included in the request.
func (c *InterconnectAttachmentsClient) Insert(ctx context.Context, req *computepb.InsertInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of interconnect attachments contained within the specified region.
func (c *InterconnectAttachmentsClient) List(ctx context.Context, req *computepb.ListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified interconnect attachment with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *InterconnectAttachmentsClient) Patch(ctx context.Context, req *computepb.PatchInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetLabels sets the labels on an InterconnectAttachment. To learn more about labels, read the Labeling Resources documentation.
func (c *InterconnectAttachmentsClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type interconnectAttachmentsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *RegionOperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing InterconnectAttachmentsClient
	CallOptions **InterconnectAttachmentsCallOptions
}

// NewInterconnectAttachmentsRESTClient creates a new interconnect attachments rest client.
//
// The InterconnectAttachments API.
func NewInterconnectAttachmentsRESTClient(ctx context.Context, opts ...option.ClientOption) (*InterconnectAttachmentsClient, error) {
	clientOpts := append(defaultInterconnectAttachmentsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultInterconnectAttachmentsRESTCallOptions()
	c := &interconnectAttachmentsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewRegionOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &InterconnectAttachmentsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultInterconnectAttachmentsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *interconnectAttachmentsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *interconnectAttachmentsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *interconnectAttachmentsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves an aggregated list of interconnect attachments.
func (c *interconnectAttachmentsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator {
	it := &InterconnectAttachmentsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListInterconnectAttachmentsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]InterconnectAttachmentsScopedListPair, string, error) {
		resp := &computepb.InterconnectAttachmentAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/interconnectAttachments", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp

		elems := make([]InterconnectAttachmentsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, InterconnectAttachmentsScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Delete deletes the specified interconnect attachment.
func (c *interconnectAttachmentsRESTClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments/%v", req.GetProject(), req.GetRegion(), req.GetInterconnectAttachment())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Get returns the specified interconnect attachment.
func (c *interconnectAttachmentsRESTClient) Get(ctx context.Context, req *computepb.GetInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.InterconnectAttachment, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments/%v", req.GetProject(), req.GetRegion(), req.GetInterconnectAttachment())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.InterconnectAttachment{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates an InterconnectAttachment in the specified project using the data included in the request.
func (c *interconnectAttachmentsRESTClient) Insert(ctx context.Context, req *computepb.InsertInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInterconnectAttachmentResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req != nil && req.ValidateOnly != nil {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// List retrieves the list of interconnect attachments contained within the specified region.
func (c *interconnectAttachmentsRESTClient) List(ctx context.Context, req *computepb.ListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentIterator {
	it := &InterconnectAttachmentIterator{}
	req = proto.Clone(req).(*computepb.ListInterconnectAttachmentsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InterconnectAttachment, string, error) {
		resp := &computepb.InterconnectAttachmentList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch updates the specified interconnect attachment with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *interconnectAttachmentsRESTClient) Patch(ctx context.Context, req *computepb.PatchInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInterconnectAttachmentResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments/%v", req.GetProject(), req.GetRegion(), req.GetInterconnectAttachment())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// SetLabels sets the labels on an InterconnectAttachment. To learn more about labels, read the Labeling Resources documentation.
func (c *interconnectAttachmentsRESTClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectAttachmentRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionSetLabelsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/interconnectAttachments/%v/setLabels", req.GetProject(), req.GetRegion(), req.GetResource())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// InterconnectAttachmentIterator manages a stream of *computepb.InterconnectAttachment.
type InterconnectAttachmentIterator struct {
	items    []*computepb.InterconnectAttachment
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.InterconnectAttachment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InterconnectAttachmentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InterconnectAttachmentIterator) Next() (*computepb.InterconnectAttachment, error) {
	var item *computepb.InterconnectAttachment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InterconnectAttachmentIterator) bufLen() int {
	return len(it.items)
}

func (it *InterconnectAttachmentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// InterconnectAttachmentsScopedListPair is a holder type for string/*computepb.InterconnectAttachmentsScopedList map entries
type InterconnectAttachmentsScopedListPair struct {
	Key   string
	Value *computepb.InterconnectAttachmentsScopedList
}

// InterconnectAttachmentsScopedListPairIterator manages a stream of InterconnectAttachmentsScopedListPair.
type InterconnectAttachmentsScopedListPairIterator struct {
	items    []InterconnectAttachmentsScopedListPair
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []InterconnectAttachmentsScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InterconnectAttachmentsScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InterconnectAttachmentsScopedListPairIterator) Next() (InterconnectAttachmentsScopedListPair, error) {
	var item InterconnectAttachmentsScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InterconnectAttachmentsScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *InterconnectAttachmentsScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
