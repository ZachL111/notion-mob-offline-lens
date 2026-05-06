# Field Notes

The fixture is small on purpose, which makes each domain case carry real weight.

The domain cases cover `form pressure`, `sync drift`, `local state`, and `conflict cost`. They sit beside the smaller starter fixture so the project has both a compact scoring check and a domain-flavored review check.

`recovery` tells me the happy path still works. `baseline` tells me whether the guardrail still has teeth.

The language-specific addition keeps the review model as a small package with table tests.
