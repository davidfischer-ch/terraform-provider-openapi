# `x-timeout`


## Properties

* **Scope**: operation (POST, GET, PUT/PATCH, DELETE) per resource
* **Value**: duration string e.g. `"30m"` or `"10s"`


## Description

Sets the recommended default timeout for the corresponding Terraform action (create, read,
update, delete). Each HTTP operation carries its own value. When absent, a 20-minute default
applies (Terraform's standard default resource timeout).


### Timeout resolution (priority order)

1. User value in the resource `timeouts` block (highest priority)
2. `x-timeout` value in the OAS spec for that operation
3. 20-minute provider fallback (Terraform's standard default)


### How each operation uses its timeout

**read**: the `get` OAS3 operation x-timeout bounds all individual GET requests and the whole
operation. Including the Read CRUD operation, any listing done by a data source, and each polling
GET fired during any async operation described below.

**create**: the `post` OAS3 operation x-timeout bounds the single HTTP POST request and the whole
create operation.

**update**: the `patch`/`put` OAS3 operation x-timeout bounds the PATCH/PUT request and the whole
update operation. PATCH takes priority when both are declared; the x-timeout is read from whichever
method is actually used.

**delete**: the `delete` OAS3 operation x-timeout bounds the DELETE HTTP request and the whole
delete operation. Each individual polling GET also applies the `get` OAS3 operation x-timeout as a
nested sub-deadline, so the effective per-GET limit is the minimum of the remaining delete budget
and the read timeout.

Notes:

All timeouts are enforced as `context.WithTimeout` deadlines. The delete conflict-retry and polling
loops check `ctx.Done()` at each sleep so the deadline is respected promptly.

The resolved timeout is applied to the whole CRUD operation, not just the individual HTTP request.
If the operation does not complete within the deadline, Terraform reports a context deadline
exceeded error and marks the operation as failed.


### User overrides

Users can override any action via the `timeouts` block in the resource declaration:

```hcl
resource "openapi_vm" "my_vm" {
  timeouts {
    create = "60m"
    read   = "10s"
    update = "20m"
    delete = "15m"
  }
}
```

Timeout values set by the user are persisted to state and restored on subsequent operations.
When the `timeouts` block is absent, the spec defaults (or the 20-minute fallback) apply and
nothing is stored in state.


## Example

```yaml
/vms/:
  post:
    x-timeout: "30m"    # create: bound the POST request
/vms/{id}:
  get:
    x-timeout: "10s"    # read: bound every GET (Read, data source list, delete polling)
  put:
    x-timeout: "15m"    # update: bound the PUT/PATCH request
  delete:
    x-timeout: "10m"    # delete: bound the DELETE call + all polling until gone
```


## Prior art

No equivalent found. dikhan `x-terraform-resource-timeout` is a coarser single-value variant
(per resource, not per action); the per-action design here is a new improvement.
