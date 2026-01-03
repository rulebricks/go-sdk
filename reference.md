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
        UserGroups: []string{
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

**userGroups:** `[]string` — List of user group names or IDs to assign to the user. All specified groups must exist in your organization.
    
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

**userGroups:** `[]string` — List of user group names or IDs to assign to the user.
    
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

<details><summary><code>client.Assets.Import(request) -> *sdk.ImportManifestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import rules, flows, contexts, and values from an RBM manifest file.
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
request := &sdk.ImportManifestRequest{
        Manifest: &sdk.ImportManifestRequestManifest{
            Version: sdk.String(
                "1.0",
            ),
            Rules: []map[string]any{
                map[string]any{
                    "name": "Pricing Rule",
                    "slug": "pricing-rule",
                },
            },
            Flows: []map[string]any{
                map[string]any{
                    "name": "Onboarding Flow",
                    "slug": "onboarding-flow",
                },
            },
            Contexts: []map[string]any{
                map[string]any{
                    "name": "Customer",
                    "slug": "customer",
                },
            },
            Values: []map[string]any{
                map[string]any{
                    "key": "tax_rate",
                    "value": 0.08,
                },
            },
        },
        Overwrite: sdk.Bool(
            false,
        ),
    }
client.Assets.Import(
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

**manifest:** `*sdk.ImportManifestRequestManifest` — The RBM manifest object containing assets to import.
    
</dd>
</dl>

<dl>
<dd>

**overwrite:** `*bool` — Whether to overwrite existing assets with the same ID/slug.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Export(request) -> *sdk.ExportAssetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export selected rules, flows, contexts, and values to an RBM manifest file.
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
request := &sdk.ExportManifestRequest{
        Rules: []string{
            "pricing-rule",
            "eligibility-check",
        },
        Flows: []string{
            "onboarding-flow",
        },
        Contexts: []string{
            "customer",
        },
        Values: []string{
            "tax_rate",
            "discount_threshold",
        },
    }
client.Assets.Export(
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

**rules:** `[]string` — Rule IDs or slugs to export.
    
</dd>
</dl>

<dl>
<dd>

**flows:** `[]string` — Flow IDs or slugs to export.
    
</dd>
</dl>

<dl>
<dd>

**contexts:** `[]string` — Context IDs or slugs to export.
    
</dd>
</dl>

<dl>
<dd>

**values:** `[]string` — Value IDs or names to export.
    
</dd>
</dl>

<dl>
<dd>

**includeAll:** `*bool` — Export all assets of specified types.
    
</dd>
</dl>

<dl>
<dd>

**preview:** `*bool` — Return a preview of what would be exported without the full data.
    
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
        UserGroups: []string{
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

**userGroups:** `[]string` — Optional array of access group names or IDs. If omitted and user belongs to access groups, values will be assigned to all user's access groups. Required if values should be restricted to specific access groups.
    
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

## Contexts
<details><summary><code>client.Contexts.GetInstance(Slug, Instance) -> *sdk.ContextInstanceState</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the current state of a context instance.
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
request := &sdk.GetInstanceContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.GetInstance(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Submit(Slug, Instance, request) -> *sdk.SubmitContextDataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit data to a context instance, creating it if it doesn't exist. May trigger bound rule/flow evaluations.
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
request := &sdk.SubmitContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
        Body: map[string]any{
            "email": "customer@example.com",
            "age": 30,
        },
    }
client.Contexts.Submit(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>

<dl>
<dd>

**request:** `sdk.SubmitContextDataRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.DeleteInstance(Slug, Instance) -> *sdk.DeleteContextInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific context instance and its history.
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
request := &sdk.DeleteInstanceContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.DeleteInstance(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.GetHistory(Slug, Instance) -> *sdk.ContextInstanceHistory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the change history for a context instance.
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
request := &sdk.GetHistoryContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.GetHistory(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>

<dl>
<dd>

**field:** `*string` — Filter history to a specific field.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of history entries to return.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.GetPending(Slug, Instance) -> *sdk.ContextInstancePendingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of rules/flows that need to be evaluated for this instance.
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
request := &sdk.GetPendingContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.GetPending(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Solve(Slug, Instance, RuleSlug, request) -> *sdk.SolveContextRuleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a specific rule using the context instance's state as input.
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
request := &sdk.SolveContextRuleRequest{
        Slug: "customer",
        Instance: "cust-12345",
        RuleSlug: "eligibility-check",
    }
client.Contexts.Solve(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>

<dl>
<dd>

**ruleSlug:** `string` — The unique slug for the rule.
    
</dd>
</dl>

<dl>
<dd>

**additionalData:** `map[string]any` — Additional data to merge with instance state for rule evaluation.
    
</dd>
</dl>

<dl>
<dd>

**persist:** `*bool` — Whether to persist derived outputs to the instance.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Cascade(Slug, Instance, request) -> *sdk.CascadeContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Trigger re-evaluation of all bound rules and flows for the instance.
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
request := &sdk.CascadeContextRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.Cascade(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>

<dl>
<dd>

**maxDepth:** `*int` — Maximum depth for cascading evaluations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Execute(Slug, Instance, FlowSlug, request) -> *sdk.SolveContextFlowResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a specific flow using the context instance's state as input.
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
request := &sdk.SolveContextFlowRequest{
        Slug: "customer",
        Instance: "cust-12345",
        FlowSlug: "onboarding-flow",
    }
client.Contexts.Execute(
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

**slug:** `string` — The unique slug for the context.
    
</dd>
</dl>

<dl>
<dd>

**instance:** `string` — The unique identifier for the context instance.
    
</dd>
</dl>

<dl>
<dd>

**flowSlug:** `string` — The unique slug for the flow.
    
</dd>
</dl>

<dl>
<dd>

**additionalData:** `map[string]any` — Additional data to merge with instance state for flow execution.
    
</dd>
</dl>

<dl>
<dd>

**persist:** `*bool` — Whether to persist derived outputs to the instance.
    
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

## Contexts Admin
<details><summary><code>client.Contexts.Admin.List() -> sdk.ContextListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all contexts (entities) for the authenticated user.
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
client.Contexts.Admin.List(
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

<details><summary><code>client.Contexts.Admin.Create(request) -> sdk.CreateContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new context (entity) for the authenticated user.
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
request := &contexts.CreateContextRequest{
        Name: "Customer",
        Description: sdk.String(
            "Represents a customer in the system",
        ),
        Schema: []*contexts.CreateContextRequestSchemaItem{
            &contexts.CreateContextRequestSchemaItem{
                Key: sdk.String(
                    "email",
                ),
                Name: sdk.String(
                    "Email",
                ),
                Type: sdk.String(
                    "string",
                ),
            },
            &contexts.CreateContextRequestSchemaItem{
                Key: sdk.String(
                    "age",
                ),
                Name: sdk.String(
                    "Age",
                ),
                Type: sdk.String(
                    "number",
                ),
            },
        },
    }
client.Contexts.Admin.Create(
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

**name:** `string` — The name of the context.
    
</dd>
</dl>

<dl>
<dd>

**slug:** `*string` — Optional custom slug. Auto-generated if not provided.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of the context.
    
</dd>
</dl>

<dl>
<dd>

**schema:** `[]*contexts.CreateContextRequestSchemaItem` — Initial schema fields for the context.
    
</dd>
</dl>

<dl>
<dd>

**autoExecuteDecisions:** `*bool` — When true (default), bound rules and flows automatically execute when their inputs are satisfied.
    
</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int` — Time-to-live in seconds for live context instances. Instances expire after this duration.
    
</dd>
</dl>

<dl>
<dd>

**historyLimit:** `*int` — Maximum number of history entries to retain per field.
    
</dd>
</dl>

<dl>
<dd>

**onSchemaMismatch:** `*contexts.CreateContextRequestOnSchemaMismatch` — How to handle fields that don't match the schema.
    
</dd>
</dl>

<dl>
<dd>

**webhookOnSolve:** `*string` — Webhook URL called when a rule or flow successfully solves.
    
</dd>
</dl>

<dl>
<dd>

**webhookOnExpire:** `*string` — Webhook URL called when a live context expires due to TTL.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Admin.Get(ID) -> *sdk.ContextDetail</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific context by its ID.
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
request := &contexts.GetAdminRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Contexts.Admin.Get(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Admin.Update(ID, request) -> sdk.UpdateContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing context's properties and schema.
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
request := &contexts.UpdateContextRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        Name: sdk.String(
            "Updated Customer",
        ),
        Description: sdk.String(
            "Updated description for premium customers",
        ),
    }
client.Contexts.Admin.Update(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the context.
    
</dd>
</dl>

<dl>
<dd>

**slug:** `*string` — The slug of the context.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of the context.
    
</dd>
</dl>

<dl>
<dd>

**schema:** `[]*contexts.UpdateContextRequestSchemaItem` — Updated schema fields for the context.
    
</dd>
</dl>

<dl>
<dd>

**autoExecuteDecisions:** `*bool` — When true, bound rules and flows automatically execute when their inputs are satisfied.
    
</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int` — Time-to-live in seconds for live context instances. Instances expire after this duration.
    
</dd>
</dl>

<dl>
<dd>

**historyLimit:** `*int` — Maximum number of history entries to retain per field.
    
</dd>
</dl>

<dl>
<dd>

**onSchemaMismatch:** `*contexts.UpdateContextRequestOnSchemaMismatch` — How to handle fields that don't match the schema.
    
</dd>
</dl>

<dl>
<dd>

**webhookOnSolve:** `*string` — Webhook URL called when a rule or flow successfully solves.
    
</dd>
</dl>

<dl>
<dd>

**webhookOnExpire:** `*string` — Webhook URL called when a live context expires due to TTL.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Admin.Delete(ID) -> *sdk.DeleteContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific context and all its instances.
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
request := &contexts.DeleteAdminRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Contexts.Admin.Delete(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Contexts Relationships
<details><summary><code>client.Contexts.Relationships.List(ID) -> *sdk.ContextRelationshipsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all relationships for a specific context.
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
request := &contexts.ListRelationshipsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Contexts.Relationships.List(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Relationships.Create(ID, request) -> sdk.CreateRelationshipResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new relationship between two contexts.
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
request := &contexts.CreateRelationshipRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        TargetContextID: "b2c3d4e5-f6a7-8901-bcde-f12345678901",
        Type: contexts.CreateRelationshipRequestTypeOneToMany,
        ForeignKey: "customer_id",
        Name: sdk.String(
            "Customer Orders",
        ),
    }
client.Contexts.Relationships.Create(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>

<dl>
<dd>

**targetContextID:** `string` — The ID of the target context.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*contexts.CreateRelationshipRequestType` — The type of relationship.
    
</dd>
</dl>

<dl>
<dd>

**foreignKey:** `string` — The field key to use as the foreign key.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Display name for the relationship.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of the relationship.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Contexts.Relationships.Delete(ID, Relationship) -> *sdk.DeleteRelationshipResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific relationship between contexts.
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
request := &contexts.DeleteRelationshipsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        Relationship: "c3d4e5f6-a7b8-9012-cdef-123456789012",
    }
client.Contexts.Relationships.Delete(
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

**id:** `string` — The unique identifier for the context.
    
</dd>
</dl>

<dl>
<dd>

**relationship:** `string` — The unique identifier for the relationship.
    
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
