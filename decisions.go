// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "sdk/core"
	time "time"
)

type QueryRequest struct {
	// The slug of the rule to query logs for.
	Slug string `json:"-"`
	// Start date for the query range (ISO8601 format).
	From *time.Time `json:"-"`
	// End date for the query range (ISO8601 format).
	To *time.Time `json:"-"`
	// Cursor for pagination.
	Cursor *string `json:"-"`
	// Number of results to return per page.
	Limit *int `json:"-"`
}

type QueryResponse struct {
	Data   []map[string]interface{} `json:"data,omitempty"`
	Cursor *string                  `json:"cursor,omitempty"`

	_rawJSON json.RawMessage
}

func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler QueryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*q = QueryResponse(value)
	q._rawJSON = json.RawMessage(data)
	return nil
}

func (q *QueryResponse) String() string {
	if len(q._rawJSON) > 0 {
		if value, err := core.StringifyJSON(q._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(q); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", q)
}
