// This file was auto-generated by Fern from our API Definition.

package users

type CreateUserGroupRequest struct {
	// Unique name of the user group.
	Name string `json:"name" url:"-"`
	// Description of the user group.
	Description *string `json:"description,omitempty" url:"-"`
}
