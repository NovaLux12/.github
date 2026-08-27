# NovaLux12/.github

Fleet-level GitHub config: shared reusable workflows for the NovaLux12 repos, so CI logic
lives in exactly one place and per-repo callers stay 4 lines.

## Reusable workflows

| Workflow | Purpose | Caller |
|---|---|---|
| [`go-ci.yml`](.github/workflows/go-ci.yml) | vet + test + gofmt check + build/--help smoke test | `uses: NovaLux12/.github/.github/workflows/go-ci.yml@main` |
| [`python-ci.yml`](.github/workflows/python-ci.yml) | pip install + pytest | `uses: NovaLux12/.github/.github/workflows/python-ci.yml@main` |

Inputs (all optional): `go-version` (default `1.26`), `python-version` (default `3.11`),
`working-directory` (default `.`), `run-build` (Go, default true), `install` (Python override).

## Adopting in a repo

Replace `.github/workflows/ci.yml` with:

```yaml
name: CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  ci:
    uses: NovaLux12/.github/.github/workflows/go-ci.yml@main
```

(Python repos: swap the `uses:` line. Note: the required-check context name becomes the caller's
job name, e.g. `ci` — renaming job names invalidates any existing branch-protection required
checks; none currently exist on Nova repos.)

## Self-test

[`self-test.yml`](.github/workflows/self-test.yml) runs both reusable workflows against the
`samples/` modules on every change to workflows/samples — green there means the shared
definitions work before any repo adopts them.
