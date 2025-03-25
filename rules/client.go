// This file was auto-generated by Fern from our API Definition.

package rules

import (
	context "context"
	http "net/http"
	sdk "sdk"
	core "sdk/core"
	internal "sdk/internal"
	option "sdk/option"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Executes a single rule identified by a unique slug. The request and response formats are dynamic, dependent on the rule configuration.
func (c *Client) Solve(
	ctx context.Context,
	// The unique identifier for the resource.
	slug string,
	request sdk.DynamicRequestPayload,
	opts ...option.RequestOption,
) (sdk.DynamicResponsePayload, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/solve/%v",
		slug,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response sdk.DynamicResponsePayload
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Executes a particular rule against multiple request data payloads provided in a list.
func (c *Client) BulkSolve(
	ctx context.Context,
	// The unique identifier for the resource.
	slug string,
	request []sdk.DynamicRequestPayload,
	opts ...option.RequestOption,
) ([]*sdk.BulkRuleResponseItem, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/bulk-solve/%v",
		slug,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response []*sdk.BulkRuleResponseItem
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Executes multiple rules or flows in parallel based on a provided mapping of rule/flow slugs to payloads.
func (c *Client) ParallelSolve(
	ctx context.Context,
	request sdk.ParallelSolveRequest,
	opts ...option.RequestOption,
) (sdk.ParallelSolveResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := baseURL + "/api/v1/parallel-solve"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response sdk.ParallelSolveResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
