// This file was auto-generated by Fern from our API Definition.

package rules

import (
	context "context"
	http "net/http"
	sdk "sdk"
	assets "sdk/assets"
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

// Delete a specific rule by its ID.
func (c *Client) Delete(
	ctx context.Context,
	request *assets.DeleteRuleRequest,
	opts ...option.RequestOption,
) (*sdk.SuccessMessage, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://rulebricks.com/api/v1",
	)
	endpointURL := baseURL + "/admin/rules/delete"
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
		404: func(apiError *core.APIError) error {
			return &sdk.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.SuccessMessage
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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

// Export a specific rule by its ID.
func (c *Client) Pull(
	ctx context.Context,
	request *assets.RulesPullRequest,
	opts ...option.RequestOption,
) (sdk.RuleExport, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://rulebricks.com/api/v1",
	)
	endpointURL := baseURL + "/admin/rules/export"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &sdk.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response sdk.RuleExport
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Import a rule into the user's account.
func (c *Client) Push(
	ctx context.Context,
	request *assets.ImportRuleRequest,
	opts ...option.RequestOption,
) (sdk.RuleExport, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://rulebricks.com/api/v1",
	)
	endpointURL := baseURL + "/admin/rules/import"
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
		403: func(apiError *core.APIError) error {
			return &sdk.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response sdk.RuleExport
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

// List all rules in the organization. Optionally filter by folder name or ID.
func (c *Client) List(
	ctx context.Context,
	request *assets.RulesListRequest,
	opts ...option.RequestOption,
) (sdk.RuleListResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://rulebricks.com/api/v1",
	)
	endpointURL := baseURL + "/admin/rules/list"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
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

	var response sdk.RuleListResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
