# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Project Does

MonoHook is an AWS Lambda function that bridges Monobank (Ukrainian bank) webhooks to YNAB (You Need A Budget). When a Monobank transaction occurs, the webhook triggers the Lambda, which transforms and creates a corresponding transaction in YNAB.

## Build & Test Commands

```bash
# Run all tests
go test ./...

# Run a single test
go test ./internal/cpu -run TestShortenString

# Build Lambda zip locally
make build
```

## Architecture

**Request flow:** API Gateway → Lambda handler (`cmd/main.go`) → CPU processor (`internal/cpu`) → YNAB client (`pkg/ynab`)

- `cmd/main.go` — Lambda entry point. `init()` runs once per cold start (fetches YNAB token from Secrets Manager, validates env vars). `HandleRequest` processes webhooks. GET returns 200 (health check). POST parses Monobank webhook and creates YNAB transaction. Always returns 200 to prevent webhook retries.
- `internal/cfg` — Fetches YNAB token from AWS Secrets Manager (secret name: `ynabToken`, region: `eu-central-1`, format: `{"token": "..."}`)
- `internal/cpu` — Core business logic. Transforms Monobank transactions to YNAB format: amount conversion (×10 for milliunits), description/comment parsing into memo+payee, import ID generation.
- `pkg/ynab` — HTTP client for YNAB API (`POST /budgets/{id}/transactions`)
- `pkg/monobank` — Type definitions only (no logic) for Monobank webhook payloads

## Key Details

- **Lambda runtime:** `provided.al2023` on `arm64` — build must use `CGO_ENABLED=0 GOARCH=arm64 GOOS=linux`
- **Required Lambda env vars:** `BUDGET_ID`, `ACCOUNT_ID`
- **Test patterns:** Table-driven tests with `testify/require`. YNAB client tests use `httptest.Server` for mocking.
- **CI/CD:** GitHub Actions (`.github/workflows/deploy.yml`) — tests on all PRs, deploys to Lambda on push to main
