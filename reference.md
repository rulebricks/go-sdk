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
            "age": 30,
            "email": "jdoe@acme.co",
            "name": "John Doe",
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
                "age": 30,
                "email": "jdoe@acme.co",
                "name": "John Doe",
            },
            map[string]any{
                "age": 28,
                "email": "jane@example.com",
                "name": "Jane Doe",
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
        "eligibility": &sdk.ParallelSolveRequestValue{
            Rule: sdk.String(
                "1ef03ms",
            ),
        },
        "offers": &sdk.ParallelSolveRequestValue{
            Flow: sdk.String(
                "OvmsYwn",
            ),
        },
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
            "age": 30,
            "email": "jdoe@acme.co",
            "name": "John Doe",
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

Invite a new user to the organization or update role or user group data for an existing user.
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

List all users (including the admin and all team members) in the organization with their details including email, name, API key, role, user groups, and join date.
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

<details><summary><code>client.Assets.ImportRbm(request) -> *sdk.ImportManifestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import rules, flows, contexts, and values from an Rulebricks manifest file (*.rbm).
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
            Entities: []map[string]any{
                map[string]any{
                    "name": "Customer",
                    "slug": "customer",
                },
            },
            Values: []map[string]any{
                map[string]any{
                    "name": "tax_rate",
                    "value": 0.08,
                },
            },
        },
        ConflictStrategy: sdk.ImportManifestRequestConflictStrategyUpdate.Ptr(),
    }
client.Assets.ImportRbm(
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

**manifest:** `*sdk.ImportManifestRequestManifest` — The RBM manifest object containing assets to import. Asset objects inside the manifest intentionally preserve `.rbm`/database casing so exported manifests can be imported without rewriting asset payloads.
    
</dd>
</dl>

<dl>
<dd>

**conflictStrategy:** `*sdk.ImportManifestRequestConflictStrategy` — How to handle conflicts with existing assets. 'update' overwrites, 'skip' ignores, 'error' fails.
    
</dd>
</dl>

<dl>
<dd>

**targetFolderName:** `*string` — Optional folder name to place imported assets into. Created if it doesn't exist.
    
</dd>
</dl>

<dl>
<dd>

**legacyRuleMapping:** `map[string]*sdk.ImportManifestRequestLegacyRuleMappingValue` — Optional mapping for legacy flow imports to reuse existing rules.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.ExportRbm(request) -> *sdk.ExportRbmAssetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export selected rules, flows, contexts, and values to an Rulebricks manifest file (*.rbm).
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
        RootType: sdk.ExportManifestRequestRootTypeRule,
        RootIDs: []string{
            "pricing-rule",
            "eligibility-check",
        },
        IncludeDownstream: sdk.Bool(
            false,
        ),
    }
client.Assets.ExportRbm(
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

**rootType:** `*sdk.ExportManifestRequestRootType` — The type of root asset to export. All dependencies will be included.
    
</dd>
</dl>

<dl>
<dd>

**rootIDs:** `[]string` — Array of IDs for the root assets to export. Dependencies are automatically resolved.
    
</dd>
</dl>

<dl>
<dd>

**includeDownstream:** `*bool` — For context exports, whether to include rules and flows bound to the context.
    
</dd>
</dl>

<dl>
<dd>

**manifestName:** `*string` — Optional name for the exported manifest.
    
</dd>
</dl>

<dl>
<dd>

**manifestDescription:** `*string` — Optional description for the exported manifest.
    
</dd>
</dl>

<dl>
<dd>

**previewOnly:** `*bool` — If true, returns a preview of what would be exported without the full data.
    
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
            "Age": 30,
            "Favorite Color": "blue",
            "Hobbies": []any{
                "reading",
                "cycling",
            },
            "Is Student": false,
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

**userGroups:** `[]string` — Optional array of user group names or IDs. If omitted and user belongs to user groups, values will be assigned to all user's user groups. Required if values should be restricted to specific user groups.
    
</dd>
</dl>

<dl>
<dd>

**metadataByName:** `map[string]map[string]any` — Optional metadata keyed by dynamic value name. This is the canonical snake_case field; legacy clients may still send `metadataByName`.
    
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
<details><summary><code>client.Contexts.Get(Slug, Instance) -> *sdk.ContextInstanceState</code></summary>
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
request := &sdk.GetContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.Get(
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
            "age": 30,
            "email": "customer@example.com",
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

<details><summary><code>client.Contexts.Delete(Slug, Instance) -> *sdk.DeleteContextInstanceResponse</code></summary>
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
request := &sdk.DeleteContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
    }
client.Contexts.Delete(
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
request := &sdk.SolveContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
        RuleSlug: "eligibility-check",
        Body: map[string]any{},
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

**request:** `sdk.SolveContextRuleRequest` 
    
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
request := &sdk.CascadeContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
        Body: map[string]any{},
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

**request:** `sdk.CascadeContextRequest` 
    
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
request := &sdk.ExecuteContextsRequest{
        Slug: "customer",
        Instance: "cust-12345",
        FlowSlug: "onboarding-flow",
        Body: map[string]any{},
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

**request:** `sdk.SolveContextFlowRequest` 
    
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

Export a specific rule by its ID. This response preserves the raw rule document casing (for example, `requestSchema`, `sampleRequest`, and `createdAt`) so it can round-trip through `/admin/rules/import` and `.rbm` workflows.
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

Create or update a rule. If `id` is provided, the matching rule is partially updated (all other fields optional). If `id` is omitted, a new rule is created (`id` and `slug` are auto-generated; all other fields required).
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
        Rule: &sdk.RuleImportPayload{
            Name: sdk.String(
                "Basic Pricing Rule",
            ),
            Description: sdk.String(
                "",
            ),
            CreatedAt: sdk.Time(
                sdk.MustParseDateTime(
                    "2026-02-12T01:29:23.000Z",
                ),
            ),
            UpdatedAt: sdk.Time(
                sdk.MustParseDateTime(
                    "2026-02-12T01:29:23.000Z",
                ),
            ),
            Published: sdk.Bool(
                false,
            ),
            RequestSchema: []*sdk.RuleImportSchemaField{
                &sdk.RuleImportSchemaField{
                    Key: "customer_tier",
                    Show: true,
                    Name: "Customer Tier",
                    Type: sdk.RuleImportSchemaFieldTypeString,
                },
                &sdk.RuleImportSchemaField{
                    Key: "order_total",
                    Show: true,
                    Name: "Order Total",
                    Type: sdk.RuleImportSchemaFieldTypeNumber,
                },
                &sdk.RuleImportSchemaField{
                    Key: "expedited",
                    Show: true,
                    Name: "Expedited",
                    Type: sdk.RuleImportSchemaFieldTypeBoolean,
                },
            },
            ResponseSchema: []*sdk.RuleImportSchemaField{
                &sdk.RuleImportSchemaField{
                    Key: "discount_rate",
                    Show: true,
                    Name: "Discount Rate",
                    Type: sdk.RuleImportSchemaFieldTypeNumber,
                },
                &sdk.RuleImportSchemaField{
                    Key: "approval_status",
                    Show: true,
                    Name: "Approval Status",
                    Type: sdk.RuleImportSchemaFieldTypeString,
                },
            },
            SampleRequest: map[string]any{
                "customer_tier": "STANDARD",
                "expedited": false,
                "order_total": 250,
            },
            TestRequest: map[string]any{
                "customer_tier": "STANDARD",
                "expedited": false,
                "order_total": 250,
            },
            SampleResponse: map[string]any{
                "approval_status": "standard",
                "discount_rate": 0,
            },
            Conditions: []*sdk.RuleImportConditionRow{
                &sdk.RuleImportConditionRow{
                    Request: map[string]*sdk.RuleImportRequestCell{
                        "customer_tier": &sdk.RuleImportRequestCell{
                            Op: "equals",
                            Args: []any{
                                "VIP",
                            },
                        },
                    },
                    Response: map[string]*sdk.RuleImportResponseCell{
                        "approval_status": &sdk.RuleImportResponseCell{
                            Value: "priority",
                        },
                        "discount_rate": &sdk.RuleImportResponseCell{
                            Value: 0.2,
                        },
                    },
                    Settings: &sdk.RuleImportRowSettings{
                        Enabled: true,
                        Priority: 0,
                        Schedule: []map[string]any{},
                    },
                },
                &sdk.RuleImportConditionRow{
                    Request: map[string]*sdk.RuleImportRequestCell{
                        "expedited": &sdk.RuleImportRequestCell{
                            Op: "equals",
                            Args: []any{
                                true,
                            },
                        },
                    },
                    Response: map[string]*sdk.RuleImportResponseCell{
                        "approval_status": &sdk.RuleImportResponseCell{
                            Value: "expedited",
                        },
                        "discount_rate": &sdk.RuleImportResponseCell{
                            Value: 0.05,
                        },
                    },
                    Settings: &sdk.RuleImportRowSettings{
                        Enabled: true,
                        Priority: 1,
                        Schedule: []map[string]any{},
                    },
                },
                &sdk.RuleImportConditionRow{
                    Request: map[string]*sdk.RuleImportRequestCell{},
                    Response: map[string]*sdk.RuleImportResponseCell{
                        "approval_status": &sdk.RuleImportResponseCell{
                            Value: "standard",
                        },
                        "discount_rate": &sdk.RuleImportResponseCell{
                            Value: 0,
                        },
                    },
                    Settings: &sdk.RuleImportRowSettings{
                        Enabled: true,
                        Priority: 2,
                        Schedule: []map[string]any{},
                    },
                },
            },
            History: []map[string]any{},
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

**rule:** `*sdk.RuleImportPayload` 
    
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

List all rules in the organization. Results are scoped to the API key holder's user groups. Optionally filter by folder name or ID, or by user group name or ID when the API key has access to that group.
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

<dl>
<dd>

**userGroup:** `*string` — Filter rules by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.
    
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

## Contexts Objects
<details><summary><code>client.Contexts.Objects.List() -> sdk.ContextListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all contexts for the authenticated user.
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
client.Contexts.Objects.List(
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

<details><summary><code>client.Contexts.Objects.Create(request) -> sdk.CreateContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new context for the authenticated user.
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
        IdentityFact: "email",
    }
client.Contexts.Objects.Create(
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

**schema:** `[]*contexts.CreateContextRequestSchemaItem` — Initial schema fields for the context. At least one field must be defined.
    
</dd>
</dl>

<dl>
<dd>

**identityFact:** `string` — The field key to use as the unique identifier for instances. Must be a key from the schema.
    
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

<details><summary><code>client.Contexts.Objects.Get(ID) -> *sdk.ContextDetail</code></summary>
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
request := &contexts.GetObjectsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Contexts.Objects.Get(
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

<details><summary><code>client.Contexts.Objects.Update(ID, request) -> *sdk.UpdateContextResponse</code></summary>
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
client.Contexts.Objects.Update(
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

<details><summary><code>client.Contexts.Objects.Delete(ID) -> *sdk.DeleteContextResponse</code></summary>
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
request := &contexts.DeleteObjectsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Contexts.Objects.Delete(
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
        ToContextID: "b2c3d4e5-f6a7-8901-bcde-f12345678901",
        RelationType: contexts.CreateRelationshipRequestRelationTypeHasMany,
        ForeignKeyFact: "customer_id",
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

**toContextID:** `string` — The ID of the target context.
    
</dd>
</dl>

<dl>
<dd>

**relationType:** `*contexts.CreateRelationshipRequestRelationType` — The type of relationship.
    
</dd>
</dl>

<dl>
<dd>

**foreignKeyFact:** `string` — The field key to use as the foreign key.
    
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

