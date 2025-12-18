# Reference
## Rules
<details><summary><code>client.Rules.Solve(Slug, request) -> sdk.DynamicResponsePayload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes a single rule identified by a unique slug. The request and response formats are dynamic, dependent on the rule configuration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SolveRulesRequest{
        Slug: "slug",
        Body: map[string]any{
            "body": map[string]any{
                "name": "Alice Johnson",
                "age": 28,
                "email": "alice.johnson@example.com",
            },
        },
    }
client.Rules.Solve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**request:** `sdk.DynamicRequestPayload` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.BulkSolve(Slug, request) -> []*sdk.BulkRuleResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes a particular rule against multiple request data payloads provided in a list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkSolveRulesRequest{
        Slug: "slug",
        Body: []sdk.DynamicRequestPayload{
            map[string]any{
                "name": "John Doe",
                "age": 30,
                "email": "jdoe@acme.co",
            },
            map[string]any{
                "name": "Jane Doe",
                "age": 28,
                "email": "jane@example.com",
            },
        },
    }
client.Rules.BulkSolve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]sdk.DynamicRequestPayload` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.ParallelSolve(request) -> sdk.ParallelSolveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes multiple rules or flows in parallel based on a provided mapping of rule/flow slugs to payloads.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]*sdk.ParallelSolveRequestValue{
        "body": &sdk.ParallelSolveRequestValue{},
    }
client.Rules.ParallelSolve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `sdk.ParallelSolveRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Flows
<details><summary><code>client.Flows.Execute(Slug, request) -> sdk.DynamicResponsePayload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a flow by its slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ExecuteFlowsRequest{
        Slug: "slug",
        Body: map[string]any{
            "body": map[string]any{
                "name": "Alice Johnson",
                "age": 28,
                "email": "alice.johnson@example.com",
            },
        },
    }
client.Flows.Execute(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**request:** `sdk.DynamicRequestPayload` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Decisions
<details><summary><code>client.Decisions.Query() -> *sdk.DecisionLogResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query decision logs with support for the decision data query language, rule/status filters, date ranges, and pagination. The query language supports field comparisons (e.g., `alpha=0`, `score>10`), contains/not-contains (e.g., `name:John`, `status!:error`), boolean logic (`AND`, `OR`), and parentheses for grouping.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.QueryDecisionsRequest{
        Search: sdk.String(
            "status=200",
        ),
        Rules: sdk.String(
            "Lead Qualification,Pricing Calculator",
        ),
        Statuses: sdk.String(
            "200,400,500",
        ),
    }
client.Decisions.Query(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**search:** `*string` — Decision data query language expression to filter logs by request/response data. Supports field comparisons (`field=value`, `field>10`), contains (`field:text`), not-contains (`field!:text`), boolean operators (`AND`, `OR`), and parentheses.
    
</dd>
</dl>

<dl>
<dd>

**rules:** `*string` — Comma-separated list of rule names to filter logs by.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `*string` — Comma-separated list of HTTP status codes to filter logs by.
    
</dd>
</dl>

<dl>
<dd>

**start:** `*time.Time` — Start date for the query range (ISO8601 format).
    
</dd>
</dl>

<dl>
<dd>

**end:** `*time.Time` — End date for the query range (ISO8601 format).
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Cursor for pagination (returned from previous query).
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return per page (default: 100).
    
</dd>
</dl>

<dl>
<dd>

**count:** `*sdk.QueryDecisionsRequestCount` — If set to 'true', returns only the count of matching logs instead of the log data.
    
</dd>
</dl>

<dl>
<dd>

**slug:** `*string` — (Deprecated) Legacy parameter for filtering by rule slug. Use 'rules' parameter instead.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.Invite(request) -> *sdk.UserInviteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Invite a new user to the organization or update role or access group data for an existing user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.UserInviteRequest{
        Email: "newuser@example.com",
        Role: sdk.UserInviteRequestRoleDeveloper.Ptr(),
        AccessGroups: []string{
            "group1",
            "group2",
        },
    }
client.Users.Invite(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` — Email of the user to invite.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*sdk.UserInviteRequestRole` — System or custom role ID to assign to the user. Available system roles include 'admin', 'editor', and 'developer'.
    
</dd>
</dl>

<dl>
<dd>

**accessGroups:** `[]string` — List of access group names or IDs to assign to the user. All specified groups must exist in your organization.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.List() -> sdk.UserListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all users (including the admin and all team members) in the organization with their details including email, name, API key, role, access groups, and join date.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Create(request) -> *sdk.CreateUserResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new user directly with a password, bypassing the email invitation flow. The user can immediately log in with the provided credentials.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CreateUserRequest{
        Email: "newuser@example.com",
        Password: "securePassword123",
    }
client.Users.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` — Email address for the new user.
    
</dd>
</dl>

<dl>
<dd>

**password:** `string` — Password for the new user (minimum 8 characters). The user can log in immediately with this password.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Display name for the user.
    
</dd>
</dl>

<dl>
<dd>

**role:** `*string` — Role to assign to the user. Defaults to 'developer' if not specified.
    
</dd>
</dl>

<dl>
<dd>

**accessGroups:** `[]string` — List of access group names or IDs to assign to the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Assets
<details><summary><code>client.Assets.GetUsage() -> *sdk.UsageStatistics</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the rule execution usage of your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Assets.GetUsage(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Values
<details><summary><code>client.Values.List() -> sdk.DynamicValueListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all dynamic values for the authenticated user. Use the 'include' parameter to control whether usage information is returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ListValuesRequest{
        Include: sdk.String(
            "usage",
        ),
    }
client.Values.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*string` — Query all dynamic values containing a specific name
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — Comma-separated list of additional data to include. Use 'usage' to include which rules reference each value.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Values.Update(request) -> sdk.DynamicValueListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update existing dynamic values or add new ones for the authenticated user. Supports both flat and nested object structures. Nested objects are automatically flattened using dot notation and keys are converted to readable format (e.g., 'user_name' becomes 'User Name', nested 'user.contact_info.email' becomes 'User.Contact Info.Email').
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.UpdateValuesRequest{
        Values: map[string]any{
            "Favorite Color": "blue",
            "Age": 30,
            "Is Student": false,
            "Hobbies": []any{
                "reading",
                "cycling",
            },
        },
        AccessGroups: []string{
            "marketing",
            "developers",
        },
    }
client.Values.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**values:** `map[string]any` — A dictionary of keys and values to update or add. Supports both flat key-value pairs and nested objects. Nested objects will be automatically flattened using dot notation with readable key names (e.g., 'user.contact_info.email' becomes 'User.Contact Info.Email').
    
</dd>
</dl>

<dl>
<dd>

**accessGroups:** `[]string` — Optional array of access group names or IDs. If omitted and user belongs to access groups, values will be assigned to all user's access groups. Required if values should be restricted to specific access groups.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Values.Delete() -> *sdk.SuccessMessage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific dynamic value for the authenticated user by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DeleteValuesRequest{
        ID: "id",
    }
client.Values.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the dynamic value to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Assets Rules
<details><summary><code>client.Assets.Rules.Delete(request) -> *sdk.SuccessMessage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific rule by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.DeleteRuleRequest{
        ID: "2855f8da-2654-4df9-8903-8f797cbfe8eb",
    }
client.Assets.Rules.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the rule to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Rules.Pull() -> sdk.RuleExport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export a specific rule by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.PullRulesRequest{
        ID: "2855f8da-2654-4df9-8903-8f797cbfe8eb",
    }
client.Assets.Rules.Pull(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the rule to export.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Rules.Push(request) -> sdk.RuleExport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import a rule into the user's account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.ImportRuleRequest{
        Rule: map[string]any{
            "name": "Imported Rule",
            "description": "A rule imported via API",
        },
    }
client.Assets.Rules.Push(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**rule:** `map[string]any` — The rule data to import.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Rules.List() -> sdk.RuleListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all rules in the organization. Optionally filter by folder name or ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.ListRulesRequest{
        Folder: sdk.String(
            "Marketing Rules",
        ),
    }
client.Assets.Rules.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**folder:** `*string` — Filter rules by folder name or folder ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Assets Flows
<details><summary><code>client.Assets.Flows.List() -> sdk.FlowListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all flows in the organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Assets.Flows.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Assets Folders
<details><summary><code>client.Assets.Folders.List() -> sdk.FolderListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all rule folders for the authenticated user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Assets.Folders.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Folders.Upsert(request) -> *sdk.Folder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new rule folder or update an existing one for the authenticated user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.UpsertFolderRequest{
        Name: "Marketing Rules",
        Description: sdk.String(
            "Rules for marketing automation workflows",
        ),
    }
client.Assets.Folders.Upsert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` — Folder ID (required for updates, omit for creation)
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Name of the folder
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of the folder
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Folders.Delete(request) -> *sdk.Folder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific rule folder for the authenticated user. This does not delete the rules within the folder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &assets.DeleteFolderRequest{
        ID: "abc123",
    }
client.Assets.Folders.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the folder to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tests Rules
<details><summary><code>client.Tests.Rules.List(Slug) -> sdk.TestListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of tests associated with the rule identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.ListRulesRequest{
        Slug: "slug",
    }
client.Tests.Rules.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tests.Rules.Create(Slug, request) -> *sdk.Test</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a new test to the test suite of a rule identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.CreateRulesRequest{
        Slug: "slug",
        Body: &sdk.CreateTestRequest{
            Name: "Test 3",
            Request: map[string]any{
                "param1": "value1",
            },
            Response: map[string]any{
                "status": "success",
            },
            Critical: true,
        },
    }
client.Tests.Rules.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*sdk.CreateTestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tests.Rules.Delete(Slug, TestID) -> *sdk.Test</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a test from the test suite of a rule identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.DeleteRulesRequest{
        Slug: "slug",
        TestID: "testId",
    }
client.Tests.Rules.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**testID:** `string` — The ID of the test.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tests Flows
<details><summary><code>client.Tests.Flows.List(Slug) -> sdk.TestListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of tests associated with the flow identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.ListFlowsRequest{
        Slug: "slug",
    }
client.Tests.Flows.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tests.Flows.Create(Slug, request) -> *sdk.Test</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds a new test to the test suite of a flow identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.CreateFlowsRequest{
        Slug: "slug",
        Body: &sdk.CreateTestRequest{
            Name: "Test 3",
            Request: map[string]any{
                "param1": "value1",
            },
            Response: map[string]any{
                "status": "success",
            },
            Critical: true,
        },
    }
client.Tests.Flows.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*sdk.CreateTestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tests.Flows.Delete(Slug, TestID) -> *sdk.Test</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a test from the test suite of a flow identified by the slug.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tests.DeleteFlowsRequest{
        Slug: "slug",
        TestID: "testId",
    }
client.Tests.Flows.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — The unique identifier for the resource.
    
</dd>
</dl>

<dl>
<dd>

**testID:** `string` — The ID of the test.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Groups
<details><summary><code>client.Users.Groups.List() -> sdk.UserGroupListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all user groups available in your Rulebricks organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Groups.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Groups.Create(request) -> *sdk.UserGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new user group in your Rulebricks organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreateUserGroupRequest{
        Name: "NewGroup",
        Description: sdk.String(
            "Description of the new group.",
        ),
    }
client.Users.Groups.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Unique name of the user group.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of the user group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
