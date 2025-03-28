// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "sdk/internal"
)

// Individual response item from a bulk rule execution
type BulkRuleResponseItem struct {
	DynamicResponsePayload    DynamicResponsePayload
	BulkRuleResponseItemError *BulkRuleResponseItemError

	typ string
}

func NewBulkRuleResponseItemFromDynamicResponsePayload(value DynamicResponsePayload) *BulkRuleResponseItem {
	return &BulkRuleResponseItem{typ: "DynamicResponsePayload", DynamicResponsePayload: value}
}

func NewBulkRuleResponseItemFromBulkRuleResponseItemError(value *BulkRuleResponseItemError) *BulkRuleResponseItem {
	return &BulkRuleResponseItem{typ: "BulkRuleResponseItemError", BulkRuleResponseItemError: value}
}

func (b *BulkRuleResponseItem) GetDynamicResponsePayload() DynamicResponsePayload {
	if b == nil {
		return nil
	}
	return b.DynamicResponsePayload
}

func (b *BulkRuleResponseItem) GetBulkRuleResponseItemError() *BulkRuleResponseItemError {
	if b == nil {
		return nil
	}
	return b.BulkRuleResponseItemError
}

func (b *BulkRuleResponseItem) UnmarshalJSON(data []byte) error {
	var valueDynamicResponsePayload DynamicResponsePayload
	if err := json.Unmarshal(data, &valueDynamicResponsePayload); err == nil {
		b.typ = "DynamicResponsePayload"
		b.DynamicResponsePayload = valueDynamicResponsePayload
		return nil
	}
	valueBulkRuleResponseItemError := new(BulkRuleResponseItemError)
	if err := json.Unmarshal(data, &valueBulkRuleResponseItemError); err == nil {
		b.typ = "BulkRuleResponseItemError"
		b.BulkRuleResponseItemError = valueBulkRuleResponseItemError
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BulkRuleResponseItem) MarshalJSON() ([]byte, error) {
	if b.typ == "DynamicResponsePayload" || b.DynamicResponsePayload != nil {
		return json.Marshal(b.DynamicResponsePayload)
	}
	if b.typ == "BulkRuleResponseItemError" || b.BulkRuleResponseItemError != nil {
		return json.Marshal(b.BulkRuleResponseItemError)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BulkRuleResponseItemVisitor interface {
	VisitDynamicResponsePayload(DynamicResponsePayload) error
	VisitBulkRuleResponseItemError(*BulkRuleResponseItemError) error
}

func (b *BulkRuleResponseItem) Accept(visitor BulkRuleResponseItemVisitor) error {
	if b.typ == "DynamicResponsePayload" || b.DynamicResponsePayload != nil {
		return visitor.VisitDynamicResponsePayload(b.DynamicResponsePayload)
	}
	if b.typ == "BulkRuleResponseItemError" || b.BulkRuleResponseItemError != nil {
		return visitor.VisitBulkRuleResponseItemError(b.BulkRuleResponseItemError)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BulkRuleResponseItemError struct {
	// Error message if this specific item failed to process
	Error *string `json:"error,omitempty" url:"error,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BulkRuleResponseItemError) GetError() *string {
	if b == nil {
		return nil
	}
	return b.Error
}

func (b *BulkRuleResponseItemError) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BulkRuleResponseItemError) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkRuleResponseItemError
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkRuleResponseItemError(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkRuleResponseItemError) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// A request containing multiple rule/flow executions to be run in parallel.
type ParallelSolveRequest = map[string]*ParallelSolveRequestValue

type ParallelSolveRequestValue struct {
	// Slug of the rule to execute
	Rule *string `json:"$rule,omitempty" url:"$rule,omitempty"`
	// Slug of the flow to execute
	Flow *string `json:"$flow,omitempty" url:"$flow,omitempty"`

	ExtraProperties map[string]interface{} `json:"-" url:"-"`

	rawJSON json.RawMessage
}

func (p *ParallelSolveRequestValue) GetRule() *string {
	if p == nil {
		return nil
	}
	return p.Rule
}

func (p *ParallelSolveRequestValue) GetFlow() *string {
	if p == nil {
		return nil
	}
	return p.Flow
}

func (p *ParallelSolveRequestValue) GetExtraProperties() map[string]interface{} {
	return p.ExtraProperties
}

func (p *ParallelSolveRequestValue) UnmarshalJSON(data []byte) error {
	type embed ParallelSolveRequestValue
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = ParallelSolveRequestValue(unmarshaler.embed)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.ExtraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *ParallelSolveRequestValue) MarshalJSON() ([]byte, error) {
	type embed ParallelSolveRequestValue
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return internal.MarshalJSONWithExtraProperties(marshaler, p.ExtraProperties)
}

func (p *ParallelSolveRequestValue) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// Response from parallel rule/flow execution with results from each execution.
type ParallelSolveResponse = map[string]DynamicResponsePayload
