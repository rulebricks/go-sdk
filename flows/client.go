// This file was auto-generated by Fern from our API Definition.

package flows

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

// Execute a flow by its slug.
func (c *Client) Execute(
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
		"https://rulebricks.com/api/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/flows/%v",
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
