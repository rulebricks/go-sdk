// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "sdk/core"
	time "time"
)

type DeleteFolderRequest struct {
	// ID of the folder to delete
	Id string `json:"id"`
}

type DeleteRuleRequest struct {
	// The ID of the rule to delete.
	Id *string `json:"id,omitempty"`
}

type ExportRuleRequest struct {
	// The ID of the rule to export.
	Id string `json:"-"`
}

type ImportRuleRequest struct {
	Id             string                 `json:"id"`
	CreatedAt      time.Time              `json:"createdAt"`
	Slug           string                 `json:"slug"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	TestRequest    map[string]interface{} `json:"testRequest,omitempty"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	RequestSchema  []interface{}          `json:"requestSchema,omitempty"`
	ResponseSchema []interface{}          `json:"responseSchema,omitempty"`
	SampleRequest  map[string]interface{} `json:"sampleRequest,omitempty"`
	SampleResponse map[string]interface{} `json:"sampleResponse,omitempty"`
	Conditions     []interface{}          `json:"conditions,omitempty"`
	Published      bool                   `json:"published"`
	History        []interface{}          `json:"history,omitempty"`
}

type ListRulesRequest struct {
	// Filter rules by folder name or folder ID
	Folder *string `json:"-"`
}

type DeleteFolderResponse struct {
	// ID of the deleted folder
	Id *string `json:"id,omitempty"`
	// Name of the deleted folder
	Name *string `json:"name,omitempty"`
	// Description of the deleted folder
	Description *string `json:"description,omitempty"`
	// Last update timestamp of the deleted folder
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DeleteFolderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteFolderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteFolderResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteFolderResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeleteRuleResponse struct {
	Message *string `json:"message,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DeleteRuleResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteRuleResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteRuleResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteRuleResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type ImportRuleResponse struct {
	Name *string `json:"name,omitempty"`
	Id   *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`

	_rawJSON json.RawMessage
}

func (i *ImportRuleResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ImportRuleResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = ImportRuleResponse(value)
	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *ImportRuleResponse) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type ListFoldersResponseItem struct {
	// Unique identifier for the folder.
	Id *string `json:"id,omitempty"`
	// Name of the folder.
	Name *string `json:"name,omitempty"`
	// Description of the folder.
	Description *string `json:"description,omitempty"`
	// Timestamp of when the folder was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListFoldersResponseItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFoldersResponseItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFoldersResponseItem(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFoldersResponseItem) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListRulesResponseItem struct {
	// The unique identifier for the rule.
	Id *string `json:"id,omitempty"`
	// The date this rule was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The name of the rule.
	Name *string `json:"name,omitempty"`
	// The description of the rule.
	Description *string `json:"description,omitempty"`
	// The unique slug for the rule used in API requests.
	Slug *string `json:"slug,omitempty"`
	// The folder containing this rule
	Folder *ListRulesResponseItemFolder `json:"folder,omitempty"`
	// The published request schema for the rule.
	RequestSchema map[string]interface{} `json:"request_schema,omitempty"`
	// The published response schema for the rule.
	ResponseSchema map[string]interface{} `json:"response_schema,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListRulesResponseItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListRulesResponseItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListRulesResponseItem(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListRulesResponseItem) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The folder containing this rule
type ListRulesResponseItemFolder struct {
	// Unique identifier for the folder.
	Id *string `json:"id,omitempty"`
	// Name of the folder.
	Name *string `json:"name,omitempty"`
	// Description of the folder.
	Description *string `json:"description,omitempty"`
	// Timestamp of when the folder was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListRulesResponseItemFolder) UnmarshalJSON(data []byte) error {
	type unmarshaler ListRulesResponseItemFolder
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListRulesResponseItemFolder(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListRulesResponseItemFolder) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type UpsertFolderResponse struct {
	// ID of the created or updated folder
	Id *string `json:"id,omitempty"`
	// Name of the folder
	Name *string `json:"name,omitempty"`
	// Description of the folder
	Description *string `json:"description,omitempty"`
	// Timestamp of when the folder was updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UpsertFolderResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertFolderResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertFolderResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertFolderResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UsageResponse struct {
	// The current plan of the organization.
	Plan *string `json:"plan,omitempty"`
	// The start date of the current monthly period.
	MonthlyPeriodStart *string `json:"monthly_period_start,omitempty"`
	// The end date of the current monthly period.
	MonthlyPeriodEnd *string `json:"monthly_period_end,omitempty"`
	// The number of rule executions used this month.
	MonthlyExecutionsUsage *float64 `json:"monthly_executions_usage,omitempty"`
	// The total number of rule executions allowed this month.
	MonthlyExecutionsLimit *float64 `json:"monthly_executions_limit,omitempty"`
	// The number of rule executions remaining this month.
	MonthlyExecutionsRemaining *float64 `json:"monthly_executions_remaining,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UsageResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UsageResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UsageResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UsageResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpsertFolderRequest struct {
	// Folder ID (required for updates, omit for creation)
	Id *string `json:"id,omitempty"`
	// Name of the folder
	Name string `json:"name"`
	// Description of the folder
	Description *string `json:"description,omitempty"`
}
