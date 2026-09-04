# Reference
## Rules
<details><summary><code>client.Rules.Solve(Slug, Version, request) -> sdk.DynamicResponsePayload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes a single rule identified by a unique slug. The request and response formats are dynamic, dependent on the rule configuration. Optionally target a specific published version (e.g. `3`) or a release environment (e.g. `production`) via the `version` path segment; `latest` (the default) executes the current published version.
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
        Version: "version",
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

**version:** `string` — The version of the resource to target: a published version number (e.g. `3`), a release environment slug (e.g. `production`, always lowercase), or `latest` (default) to use the current published version.

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

<details><summary><code>client.Rules.BulkSolve(Slug, Version, request) -> []*sdk.BulkRuleResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes a particular rule against multiple request data payloads provided in a list. Optionally target a specific published version (e.g. `3`) or a release environment (e.g. `production`) via the `version` path segment; `latest` (the default) executes the current published version.
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
        Version: "version",
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

**version:** `string` — The version of the resource to target: a published version number (e.g. `3`), a release environment slug (e.g. `production`, always lowercase), or `latest` (default) to use the current published version.

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

## Infra
<details><summary><code>client.Infra.Status() -> *sdk.ScaleStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reports the fleet scale-up state. Worker counts reflect solvers that have actually joined the processing group and can accept work. Self-hosted deployments only.
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
client.Infra.Status(
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

<details><summary><code>client.Infra.Scale() -> *sdk.ScaleStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Scales up the deployment's solver fleet to its maximum capacity ahead of a known incoming batch workload. Usually takes 1-2 minutes to complete. This is completely optional, the solver fleet will scale up automatically as needed anyway. Self-hosted deployments only.
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
client.Infra.Scale(
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

## Flows
<details><summary><code>client.Flows.Execute(Slug, Version, request) -> sdk.DynamicResponsePayload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a flow by slug and optional version. Policy failures return `{ error }` with status 200, including per-item errors for bulk requests. Errors: 400 invalid input, 500 unhandled execution failure, 503 unavailable, 504 timeout.
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
        Version: "version",
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

**version:** `string` — The version of the resource to target: a published version number (e.g. `3`), a release environment slug (e.g. `production`, always lowercase), or `latest` (default) to use the current published version.

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
        Flows: sdk.String(
            "Loan Approval Flow",
        ),
        Contexts: sdk.String(
            "loans",
        ),
        Trace: sdk.String(
            "7db50259-31a0-42c1-aa3c-36409ad3c756",
        ),
        Statuses: sdk.String(
            "200,400,500",
        ),
        ItemFilter: sdk.String(
            "customer.id=cst_8f3a12",
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

**search:** `*string` — Decision data query language expression to filter logs by request/response data. Supports field comparisons (`field=value`, `field>10`), contains (`field:text`), not-contains (`field!:text`), boolean operators (`AND`, `OR`), and parentheses. A bare UUID or 32-hex term resolves as an execution/correlation-id lookup automatically.

</dd>
</dl>

<dl>
<dd>

**rules:** `*string` — Comma-separated list of rule names, IDs, or slugs to filter logs by. Names match partially; IDs and slugs match exactly.

</dd>
</dl>

<dl>
<dd>

**flows:** `*string` — Comma-separated list of flow names, IDs, or slugs to filter logs by. Matches only flow-level execution logs; the rule executions that ran inside a flow are separate records and are not included.

</dd>
</dl>

<dl>
<dd>

**contexts:** `*string` — Comma-separated list of context names or slugs to filter logs by. Matches the rule and flow executions that were triggered by those contexts (batch and interactive updates).

</dd>
</dl>

<dl>
<dd>

**trace:** `*string` — Execution-trace correlation id. Returns every decision log from one execution tree: pass a log's `decision.root_flow_execution_id` (or any `flow_execution_id` / `parallel_execution_id`, including a bulk run's per-item `item_execution_ids` entries) to retrieve the flow-level record plus all subflow and rule records from that run. On self-hosted deployments, a log's observability `trace_id` is also accepted. Combine with `rules` or `search` to narrow to a specific rule or payload within the run.

</dd>
</dl>

<dl>
<dd>

**statuses:** `*string` — Comma-separated list of HTTP status codes to filter logs by.

</dd>
</dl>

<dl>
<dd>

**includeTraces:** `*sdk.QueryDecisionsRequestIncludeTraces` — When `true`, each flow record in the response includes a decompressed `path_trace` field: the run's executed steps with their full inputs and outputs (an object for single runs, a null-aligned array matching the request array for bulk runs). Off by default - traces are stored compressed and can be large, so only enable this when you need them. Ignored in count mode.

</dd>
</dl>

<dl>
<dd>

**itemFilter:** `*string` — Bulk payload filter in the form `path=value`. For each bulk record in the results (array-shaped request/response), keeps only the items whose payload value at `path` equals `value`, slicing the `request` and `response` arrays and every index-aligned field (`decision.item_execution_ids`, `decision.item_indexes`, `decision.success_idxs`, and `path_trace` when `include_traces=true`) in lockstep so input/output alignment is preserved. Filtered records gain a `matched_items` array with the surviving items' original zero-based positions. Paths use dot notation into each item (`customer.id`, `lines.0.sku`); prefix with `request.` or `response.` to match only that side (unprefixed paths match either side). Values compare as exact scalar strings (`status=200`, `approved=true`). Non-bulk records are returned unchanged; bulk records with no matching items are returned with empty item arrays. For a bulk solve of a Collect Matches rule, each response item is itself a `{ results: [...] }` envelope, so use a positional path such as `response.results.0.status`. Wildcards are available in `search`, not `item_filter`. Typical use: combine with `search`, `flows`, or `trace` to locate a bulk run, then isolate one item's payloads and its `item_execution_ids` entry without tracking indexes. Ignored in count mode.

</dd>
</dl>

<dl>
<dd>

**start:** `*time.Time` — Start date for the query range (ISO8601 format). Hosted queries may span at most 90 days. Persistent self-hosted queries may use any range within local ClickHouse retention; PVC-less archive mode is limited to 7 days. Defaults to the applicable maximum before `end` (or before now).

</dd>
</dl>

<dl>
<dd>

**end:** `*time.Time` — End date for the query range (ISO8601 format). Defaults to now. When supplied without `start`, the query covers the preceding 90 days on hosted/table mode or 7 days in PVC-less archive mode.

</dd>
</dl>

<dl>
<dd>

**sort:** `*sdk.QueryDecisionsRequestSort` — Column to sort results by. `time` orders by execution timestamp, `name` by rule/flow name, `status` by HTTP status code, and `type` by operation (solve, bulk-solve, flows, etc.). Defaults to `time`.

</dd>
</dl>

<dl>
<dd>

**order:** `*sdk.QueryDecisionsRequestOrder` — Sort direction. Defaults to `desc`.

</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque pagination token returned by the previous response. Pass it back verbatim to fetch the next page; do not construct or modify cursor values.

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return per page (default: 100, maximum: 1000). Logs carry full request/response payloads, so use smaller limits when querying workspaces with large bulk operations. Time-sorted pagination uses a keyset cursor, so its scan cost does not grow with page depth.

</dd>
</dl>

<dl>
<dd>

**count:** `*sdk.QueryDecisionsRequestCount` — If set to 'true', returns only the count of matching logs instead of the log data.

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

<details><summary><code>client.Assets.ImportRbm(request) -> *sdk.ImportRbmAssetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import rules, flows, contexts, and values from a Rulebricks manifest file (*.rbm). Plain JSON remains supported, and clients may send the same JSON envelope gzip-compressed with `Content-Type: application/octet-stream` and `X-Rulebricks-Content-Encoding: gzip`.
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

Export selected rules, flows, contexts, and values to a Rulebricks manifest file (*.rbm). Dependencies are resolved automatically: exporting a flow includes its rules, contexts, vocabulary values, and any flows referenced by Run Flow nodes (recursively). Set `compress: true` to receive the manifest in compressed form (a compress-json array). Set `download: true` to receive that manifest directly as a streamed attachment instead of inside the `{ success, manifest }` envelope.
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

<dl>
<dd>

**compress:** `*bool` — If true, the manifest in the response is returned in compressed form: the JSON array produced by the compress-json library instead of a plain object. Compressed manifests are substantially smaller, can be saved directly as a .rbm file, and are accepted by the import endpoint as-is. Intended for raw HTTP usage and file tooling; typed SDK clients should omit this flag, since the generated response type models the manifest as an object.

</dd>
</dl>

<dl>
<dd>

**download:** `*bool` — If true, returns the manifest itself as a streamed application/json attachment with Content-Disposition, rather than the normal `{ success, manifest }` response envelope. Combine with `compress: true` for large .rbm downloads.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Values
<details><summary><code>client.Values.List() -> *sdk.ListValuesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve vocabulary values for the authenticated user. Results are scoped to the API key holder's user groups. Optionally filter by user group name or ID when the API key has access to that group. Use the 'include' parameter to control whether usage information is returned. Small workspaces may omit pagination to receive the full catalog as an array (legacy behavior); workspaces above the catalog threshold must paginate with 'limit'/'cursor', which returns { data, next_cursor, total? } ordered by name. The 'prefix' and 'type' filters narrow results to a collection or value type.
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

**name:** `*string` — Query all vocabulary values containing a specific name

</dd>
</dl>

<dl>
<dd>

**prefix:** `*string` — Only return values whose name starts with this collection prefix (e.g. 'Countries.').

</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Only return values of this type (string, number, boolean, list, date, function).

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size (default 100, max 1000). Providing limit or cursor switches the response to the paginated { data, next_cursor } envelope.

</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque pagination cursor from a previous page's next_cursor.

</dd>
</dl>

<dl>
<dd>

**userGroup:** `*string` — Filter results by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.

</dd>
</dl>

<dl>
<dd>

**include:** `*string` — Comma-separated list of additional data to include. Use 'usage' to include which rules reference each value.

</dd>
</dl>

<dl>
<dd>

**resolve:** `*bool` — By default, payloads containing value-to-value references are returned materialized (references replaced with their resolved values). Pass 'false' to return stored payloads as-is, with { "$rb": "globalValue", "id": "..." } reference markers intact, so the reference graph round-trips.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Values.Update(request) -> *sdk.UpdateValuesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update existing vocabulary values or add new ones for the authenticated user. Supports both flat and nested object structures.
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

**values:** `map[string]any` — Values to create or update. Nested objects use dot-separated names and payloads may reference other values.

</dd>
</dl>

<dl>
<dd>

**userGroups:** `[]string` — Optional array of user group names or IDs. If omitted and user belongs to user groups, values will be assigned to all user's user groups. Required if values should be restricted to specific user groups.

</dd>
</dl>

<dl>
<dd>

**metadataByName:** `map[string]map[string]any` — Optional metadata keyed by vocabulary value name. This is the canonical snake_case field; legacy clients may still send `metadataByName`. System-owned keys (managedBy, source, lockedReason, previousTokens, and archive/tombstone fields) are stripped from user payloads - managed provenance and archive state cannot be forged.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Values.Delete() -> *sdk.DeleteValueResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a value by ID. Rule and flow references block deletion; value references are replaced with the deleted value's content.
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

**id:** `string` — ID of the vocabulary value to delete

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Values.Sync(request) -> *sdk.SyncValuesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Declaratively makes a collection exactly equal to the payload. Values in the payload are upserted (Existing values keep their IDs), and values under the collection that are absent from the payload are archived by default. The `sync` endpoint supports uploading a particularly large amount of values (100k+) in chunks, using the `sync_id` parameter to track the run.
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
request := &sdk.SyncValuesRequest{
        Collection: "Medical Codes",
        Values: map[string]any{
            "A123": "A123",
            "B456": "B456",
            "C789": "C789",
        },
    }
client.Values.Sync(
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

**collection:** `string` — Collection path to sync (e.g. 'Medical Codes'). Only values under this path are affected.

</dd>
</dl>

<dl>
<dd>

**values:** `map[string]any` — Desired members of the collection, keyed relative to the collection path ('A123' becomes 'Medical Codes.A123'). Nested objects flatten with dot notation, and payloads may use ValueReference markers. An empty object empties the collection. May be omitted on a pure finalize call (sync_id + complete).

</dd>
</dl>

<dl>
<dd>

**syncID:** `*string` — Identifier for a chunked run. Repeat the call with the same sync_id for each chunk of the desired state; nothing is removed until a call with complete: true. Abandoned runs are purged after 24 hours without removing anything.

</dd>
</dl>

<dl>
<dd>

**complete:** `*bool` — Marks the run as complete, triggering the removal sweep. Implicitly true when sync_id is omitted (single-request syncs), false otherwise.

</dd>
</dl>

<dl>
<dd>

**permanentlyDelete:** `*bool` — Hard-delete removed values instead of archiving them. Removals still referenced by a rule, flow, or surviving value are archived instead and reported in 'blocked'. Self-hosted deployments retain tombstones regardless.

</dd>
</dl>

<dl>
<dd>

**dryRun:** `*bool` — Compute and return the full diff without writing anything. Only supported for single-request syncs (omit sync_id).

</dd>
</dl>

<dl>
<dd>

**userGroups:** `[]string` — Optional array of user group names to assign to written values, matching POST /values.

</dd>
</dl>

<dl>
<dd>

**metadataByName:** `map[string]map[string]any` — Optional metadata keyed by FULL value name (including the collection prefix).

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Objects
<details><summary><code>client.Objects.List() -> []*sdk.WorkspaceObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the workspace's objects (JSON Schemas). The provided API key must have permission to view vocabulary values. Results are scoped to the API key holder's user groups.
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
client.Objects.List(
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

<details><summary><code>client.Objects.Upsert(request) -> *sdk.UpsertObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates or updates an object and syncs its generated enum values.
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
request := &sdk.UpsertObjectRequest{
        Unknown: map[string]any{
            "content": map[string]any{
                "properties": map[string]any{
                    "countryCode": map[string]any{
                        "enum": []any{
                            "US",
                            "CA",
                            "GB",
                        },
                        "title": "Country Code",
                        "type": "string",
                    },
                },
                "type": "object",
            },
            "name": "Claim",
            "user_groups": []any{
                "underwriting",
            },
        },
    }
client.Objects.Upsert(
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

**request:** `*sdk.UpsertObjectRequest`

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Objects.Get(ObjectID) -> *sdk.WorkspaceObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches one object by ID or exact name. The provided API key must have permission to view vocabulary values.
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
request := &sdk.GetObjectsRequest{
        ObjectID: "objectId",
    }
client.Objects.Get(
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

**objectID:** `string` — Object ID or exact name

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Objects.Delete(ObjectID) -> *sdk.DeleteObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes the object. By default, unused values are permanently deleted while values referenced by draft, current, or historical rules, flows, or other vocabulary values are archived. Pass values=detach to keep every generated value active as an ordinary, hand-editable value.
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
request := &sdk.DeleteObjectsRequest{
        ObjectID: "objectId",
    }
client.Objects.Delete(
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

**objectID:** `string` — Object ID or exact name

</dd>
</dl>

<dl>
<dd>

**values:** `*sdk.DeleteObjectsRequestValues` — What happens to generated values: 'archive' (default) permanently deletes unused values and archives referenced values; 'detach' retains all values as active ordinary values.

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

<dl>
<dd>

**includeRelations:** `*string` — Comma-separated relationship names to include in the response under a 'relations' key (has_many relations return a list of related instance states; has_one/belongs_to return a single state or null). Use '*' for all relationships. Omitted by default - related instances are never fetched into the payload unrequested.

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

<details><summary><code>client.Contexts.Cascade(Slug, Instance, request) -> *sdk.CascadeContextResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Re-evaluate registered pending rule and flow executions for this instance after their fact or relationship dependencies may have become available. This does not run every bound asset.
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

<details><summary><code>client.Contexts.BulkIngest(Slug, request) -> *sdk.ContextBatchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit an array of records to any context in one synchronous call. Records merge into their context instances (matched by the context's identity fact), bound rules and flows whose inputs became satisfied execute, and the response returns the resolved state of every touched instance. Retries are always safe: merges are idempotent and executions are deduplicated by input hash. Fact history is recorded for tracked facts exactly as on individual writes. Clients chunk large datasets across requests.
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
request := &sdk.BulkIngestContextsRequest{
        Slug: "loan-application",
        Body: []sdk.DynamicRequestPayload{
            map[string]any{
                "amount": 12000,
                "loan_id": "APP-1",
            },
            map[string]any{
                "amount": 7300,
                "loan_id": "APP-2",
            },
        },
    }
client.Contexts.BulkIngest(
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

**include:** `*string` — Comma-separated list of per-instance fields to include in results (instance_id is always present). Omit to include everything. Valid fields: positions, is_new, status, have, need, state, expires_at, executions, executed, triggered, reason. Useful for keeping response size proportional to outcomes rather than data volume, e.g. include=status,executed.

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

<details><summary><code>client.Assets.Rules.Pull() -> *sdk.RuleExport</code></summary>
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

<details><summary><code>client.Assets.Rules.Push(request) -> *sdk.RuleExport</code></summary>
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
                        Priority: 0,
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
                        Priority: 0,
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

List all rules in the organization. Results are scoped to the API key holder's user groups. Optionally filter by folder name or ID, labels, user group name or ID when the API key has access to that group, or by name.
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

**folder:** `*string` — Filter results by folder name or folder ID.

</dd>
</dl>

<dl>
<dd>

**labels:** `*string` — Filter results to assets containing all comma-separated labels.

</dd>
</dl>

<dl>
<dd>

**userGroup:** `*string` — Filter results by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.

</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter results by name using a case-insensitive substring match.

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

List all flows in the organization. Results are scoped to the API key holder's user groups. Optionally filter by folder name or ID, labels, user group name or ID when the API key has access to that group, or by name.
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
request := &assets.ListFlowsRequest{}
client.Assets.Flows.List(
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

**folder:** `*string` — Filter results by folder name or folder ID.

</dd>
</dl>

<dl>
<dd>

**labels:** `*string` — Filter results to assets containing all comma-separated labels.

</dd>
</dl>

<dl>
<dd>

**userGroup:** `*string` — Filter results by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.

</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter results by name using a case-insensitive substring match.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Flows.Push(request) -> *sdk.FlowImportResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update a flow from the Rulebricks Flow Schema (a list of `nodes` and `connections`). The server expands the Rulebricks Flow Schema definition into the full flow graph - laying it out, wiring property/control handles, resolving referenced published rules, and backfilling node defaults - so the result both renders in the editor and executes via `/flows/{slug}` without any manual editing. If `id` is provided the matching flow is updated; otherwise a new flow is created (`id`/`slug` auto-generated). Flows auto-publish unless `_publish` is set to `false`.
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
request := &assets.ImportFlowRequest{
        Flow: &sdk.FlowImportPayload{
            Name: "Underwriting Flow",
            Nodes: []*sdk.RulebricksFlowNode{
                &sdk.RulebricksFlowNode{
                    Ref: "input",
                    Type: sdk.RulebricksFlowNodeTypeOrigin,
                    Rule: sdk.String(
                        "customer-eligibility",
                    ),
                },
                &sdk.RulebricksFlowNode{
                    Ref: "gate",
                    Type: sdk.RulebricksFlowNodeTypeContinueIf,
                    Condition: &sdk.RulebricksFlowNodeCondition{
                        Property: sdk.String(
                            "approved",
                        ),
                        Operator: sdk.String(
                            "equals",
                        ),
                        Args: []any{
                            true,
                        },
                    },
                },
                &sdk.RulebricksFlowNode{
                    Ref: "enrich",
                    Type: sdk.RulebricksFlowNodeTypeCode,
                    Outputs: []*sdk.RulebricksFlowNodeOutputsItem{
                        &sdk.RulebricksFlowNodeOutputsItem{
                            Key: "tier",
                            Type: sdk.RulebricksFlowNodeOutputsItemTypeString.Ptr(),
                        },
                    },
                    Code: sdk.String(
                        "outputs.tier = inputs.score > 700 ? 'A' : 'B'",
                    ),
                },
                &sdk.RulebricksFlowNode{
                    Ref: "out",
                    Type: sdk.RulebricksFlowNodeTypeResult,
                    Key: sdk.String(
                        "data",
                    ),
                },
            },
            Connections: []*sdk.RulebricksFlowConnection{
                &sdk.RulebricksFlowConnection{
                    From: "input",
                    To: "gate",
                    Output: sdk.String(
                        "approved",
                    ),
                },
                &sdk.RulebricksFlowConnection{
                    From: "input",
                    To: "enrich",
                    Output: sdk.String(
                        "score",
                    ),
                    Input: sdk.String(
                        "score",
                    ),
                },
                &sdk.RulebricksFlowConnection{
                    From: "gate",
                    To: "out",
                    Control: sdk.Bool(
                        true,
                    ),
                },
                &sdk.RulebricksFlowConnection{
                    From: "enrich",
                    To: "out",
                    Output: sdk.String(
                        "tier",
                    ),
                },
            },
            Publish: sdk.Bool(
                true,
            ),
        },
    }
client.Assets.Flows.Push(
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

**flow:** `*sdk.FlowImportPayload`

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Flows.Pull() -> *sdk.FlowImportPayload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export a flow into the Rulebricks Flow Schema (nodes + connections), the same shape accepted by `/admin/flows/import`. Works for flows built entirely by hand in the editor, so they can be round-tripped or version-controlled. This is distinct from the top-level `/admin/export`, which produces `.rbm` manifests.
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
request := &assets.PullFlowsRequest{}
client.Assets.Flows.Pull(
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

**id:** `*string` — The ID of the flow to export (provide `id` or `slug`).

</dd>
</dl>

<dl>
<dd>

**slug:** `*string` — The slug of the flow to export (provide `id` or `slug`).

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Flows.Delete(request) -> *sdk.SuccessMessage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific flow by its ID.
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
request := &assets.DeleteFlowRequest{
        ID: "3855f8da-2654-4df9-8903-8f797cbfe8ec",
    }
client.Assets.Flows.Delete(
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

**id:** `string` — The ID of the flow to delete.

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
request := &assets.ListFoldersRequest{}
client.Assets.Folders.List(
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

**userGroup:** `*string` — Filter results by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.

</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter results by name using a case-insensitive substring match.

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

Create a new folder or update an existing one for the authenticated user. Folders are typed to organize rules (the default), flows, or contexts.
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

<dl>
<dd>

**type_:** `*assets.UpsertFolderRequestType` — The type of assets the folder organizes. Applies on creation; ignored when updating an existing folder.

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

## Assets Contexts
<details><summary><code>client.Assets.Contexts.List() -> sdk.ContextListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all contexts for the authenticated user. Results are scoped to the API key holder's user groups. Optionally filter by folder name or ID, by user group name or ID when the API key has access to that group, or by name.
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
request := &assets.ListContextsRequest{}
client.Assets.Contexts.List(
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

**folder:** `*string` — Filter results by folder name or folder ID.

</dd>
</dl>

<dl>
<dd>

**userGroup:** `*string` — Filter results by user group name or ID. The value is validated against workspace groups. Admin/unrestricted API keys can request any group-specific view; restricted API keys may only filter to one of their assigned groups and receive a 403 when filtering outside those groups.

</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filter results by name using a case-insensitive substring match.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Contexts.Create(request) -> *sdk.CreateContextResponse</code></summary>
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
request := &assets.CreateContextRequest{
        Name: "Customer",
        Description: sdk.String(
            "Represents a customer in the system",
        ),
        Schema: &sdk.ContextSchema{
            Base: []*sdk.ContextSchemaField{
                &sdk.ContextSchemaField{
                    Key: sdk.String(
                        "email",
                    ),
                    Name: sdk.String(
                        "Email",
                    ),
                    Type: sdk.ContextSchemaFieldTypeString.Ptr(),
                    Required: sdk.Bool(
                        true,
                    ),
                },
                &sdk.ContextSchemaField{
                    Key: sdk.String(
                        "age",
                    ),
                    Name: sdk.String(
                        "Age",
                    ),
                    Type: sdk.ContextSchemaFieldTypeNumber.Ptr(),
                },
            },
            Derived: []*sdk.ContextSchemaField{},
        },
        IdentityFact: "email",
    }
client.Assets.Contexts.Create(
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

**name:** `string` — The name of the context. The context's slug is generated from it (suffixed on collision).

</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of the context.

</dd>
</dl>

<dl>
<dd>

**schema:** `*sdk.ContextSchema` — The context's schema: an object with `base` (stored facts; at least one required) and optional `derived` (expression-computed facts) field arrays.

</dd>
</dl>

<dl>
<dd>

**identityFact:** `string` — The fact key to use as the unique identifier for instances. Must be a key from schema.base.

</dd>
</dl>

<dl>
<dd>

**autoExecuteDecisions:** `*bool` — When true (default), bound rules and flows automatically execute when their inputs are satisfied.

</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int` — Time-to-live in seconds for live context instances (60 seconds to 30 days). Instances expire after this duration; each write extends the expiry.

</dd>
</dl>

<dl>
<dd>

**historyLimit:** `*int` — Maximum number of history entries to retain per field.

</dd>
</dl>

<dl>
<dd>

**onSchemaMismatch:** `*assets.CreateContextRequestOnSchemaMismatch` — How to handle submitted fields that don't match the schema: `ignore` drops them, `reject` fails the request (or the batch item), `store` persists them alongside declared facts.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Contexts.Get(ID) -> *sdk.ContextDetail</code></summary>
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
request := &assets.GetContextsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Assets.Contexts.Get(
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

<details><summary><code>client.Assets.Contexts.Update(ID, request) -> *sdk.UpdateContextResponse</code></summary>
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
request := &assets.UpdateContextRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        Name: sdk.String(
            "Updated Customer",
        ),
        Description: sdk.String(
            "Updated description for premium customers",
        ),
    }
client.Assets.Contexts.Update(
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

**name:** `*string` — The name of the context. Changing it regenerates the context's slug.

</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of the context.

</dd>
</dl>

<dl>
<dd>

**schema:** `*sdk.ContextSchema` — Updated schema for the context: an object with `base` and optional `derived` field arrays.

</dd>
</dl>

<dl>
<dd>

**identityFact:** `*string` — The fact key to use as the unique identifier for instances. Must be a key from schema.base. Caution: changing this on a context with live instances changes how future writes resolve instances.

</dd>
</dl>

<dl>
<dd>

**autoExecuteDecisions:** `*bool` — When true, bound rules and flows automatically execute when their inputs are satisfied.

</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int` — Time-to-live in seconds for live context instances (60 seconds to 30 days). Instances expire after this duration.

</dd>
</dl>

<dl>
<dd>

**historyLimit:** `*int` — Maximum number of history entries to retain per field.

</dd>
</dl>

<dl>
<dd>

**onSchemaMismatch:** `*assets.UpdateContextRequestOnSchemaMismatch` — How to handle submitted fields that don't match the schema: `ignore` drops them, `reject` fails the request (or the batch item), `store` persists them alongside declared facts.

</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Assets.Contexts.Delete(ID) -> *sdk.DeleteContextResponse</code></summary>
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
request := &assets.DeleteContextsRequest{
        ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    }
client.Assets.Contexts.Delete(
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

## Assets Contexts Relationships
<details><summary><code>client.Assets.Contexts.Relationships.List(ID) -> *sdk.ContextRelationshipsResponse</code></summary>
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
client.Assets.Contexts.Relationships.List(
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

<details><summary><code>client.Assets.Contexts.Relationships.Create(ID, request) -> sdk.CreateRelationshipResponse</code></summary>
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
            "customer_orders",
        ),
    }
client.Assets.Contexts.Relationships.Create(
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

**name:** `*string` — Optional runtime relationship key. It is normalized to lowercase snake_case; the target context slug is used when omitted.

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

<details><summary><code>client.Assets.Contexts.Relationships.Delete(ID, Relationship) -> *sdk.DeleteRelationshipResponse</code></summary>
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
client.Assets.Contexts.Relationships.Delete(
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

Adds a new test to the rule. `contains` (Contains Data, the default) finds the expected fragment anywhere in the output, `matches` (Matches Exactly) requires complete equality, and `excludes` (Excludes Data) requires the fragment to be absent.
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
            Policy: sdk.CreateTestRequestPolicyContains.Ptr(),
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

<details><summary><code>client.Tests.Rules.Run(Slug, request) -> *sdk.RunTestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes every test in the rule's test suite (or only the critical tests when `critical_only` is true) and returns a summary of which passed, which failed, and whether any CRITICAL test failed. Use the `critical_failure` flag as the signal for whether a release should be blocked. Tests always run against the latest draft of the rule; version targeting does not apply.
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
request := &tests.RunRulesRequest{
        Slug: "slug",
        Body: &sdk.RunTestsRequest{
            CriticalOnly: sdk.Bool(
                false,
            ),
        },
    }
client.Tests.Rules.Run(
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

**request:** `*sdk.RunTestsRequest`

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

Adds a new test to the flow. `contains` (Contains Data, the default) finds the expected fragment anywhere in the output, `matches` (Matches Exactly) requires complete equality, and `excludes` (Excludes Data) requires the fragment to be absent.
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
            Policy: sdk.CreateTestRequestPolicyContains.Ptr(),
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

<details><summary><code>client.Tests.Flows.Run(Slug, request) -> *sdk.RunTestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Executes every test in the flow's test suite (or only the critical tests when `critical_only` is true) against the flow's current graph and returns a summary of which passed, which failed, and whether any CRITICAL test failed. Tests always run against the latest draft of the flow; version targeting does not apply.
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
request := &tests.RunFlowsRequest{
        Slug: "slug",
        Body: &sdk.RunTestsRequest{
            CriticalOnly: sdk.Bool(
                false,
            ),
        },
    }
client.Tests.Flows.Run(
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

**request:** `*sdk.RunTestsRequest`

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

