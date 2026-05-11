# Changelog

## Release v0.1.1 (2026-05-11)

### Features

* Implement `x-timeout` extension: per-operation timeout defaults on OAS3 operations exposed
  as a `timeouts` block on each resource; values are applied as context deadlines on all HTTP calls
* Add `list` timeout read from the collection-path GET and applied to data source reads
* Validate `x-timeout` spec values at provider startup; non-positive or unparseable values
  cause an immediate error
* Validate user-supplied `timeouts` block values at plan time; values must be valid Go durations
  greater than zero (`positiveDuration` validator)


## Release v0.1.0 (2026-05-08)

Initial release.
