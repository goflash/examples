# 🚀 goflash Examples

[![CI](https://github.com/goflash/examples/actions/workflows/ci.yml/badge.svg)](https://github.com/goflash/examples/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A curated collection of small, focused examples showing how to build web apps with [goflash](https://github.com/goflash/flash).
Each example is a tiny program you can run locally to explore a specific feature or pattern.

## 📦 How to run

- Prerequisites: Go 1.23+ installed.
- Clone this repo, then run any example from its folder.

Example:

```bash
cd examples/routing/basic
go run .
```

Most examples start a server on <http://localhost:8080> unless stated otherwise.

## 📚 Examples index

| Example                                                                   | What it shows                                      |
| ------------------------------------------------------------------------- | -------------------------------------------------- |
| [basic](examples/routing/basic/)                                          | Minimal server and one route                       |
| [any_methods](examples/routing/any_methods/)                              | Handle any HTTP method on a route                  |
| [binding_json](examples/validation/binding_json/)                         | JSON request binding and JSON response             |
| [context](examples/platform/context/)                                     | Store/get values from the per-request context      |
| [context_cancel](examples/platform/context_cancel/)                       | Respect context deadlines and client disconnects   |
| [cookies](examples/io/cookies/)                                           | Setting and reading cookies                        |
| [cors_timeout](examples/middleware/cors_timeout/)                         | CORS and per-route timeout middleware              |
| [csrf](examples/middleware/csrf/)                                         | CSRF protection middleware and a simple form       |
| [custom_logger_zap](examples/platform/custom_logger_zap/)                 | Using zap via a minimal slog adapter + access logs |
| [error_handling](examples/platform/error_handling/)                       | Custom error handling and 404 responses            |
| [graceful_shutdown](examples/platform/graceful_shutdown/)                 | Graceful HTTP server shutdown                      |
| [grouping](examples/routing/grouping/)                                    | Route groups and lightweight middleware            |
| [group_nested](examples/routing/group_nested/)                            | Nested groups with layered middleware              |
| [gzip_requestid_ratelimit](examples/middleware/gzip_requestid_ratelimit/) | Gzip + request ID + rate limit + buffer middleware |
| [headers_status](examples/routing/headers_status/)                        | Setting headers and custom HTTP status             |
| [http_mount](examples/routing/http_mount/)                                | Mount net/http handlers on flash and mix routing   |
| [logger_from_context](examples/middleware/logger_from_context/)           | Structured logging with slog pulled from context   |
| [method_not_allowed](examples/routing/method_not_allowed/)                | Custom 405 Method Not Allowed handler              |
| [middleware](examples/middleware/middleware/)                             | Tiny middleware example adding a header            |
| [nethttp_compat](examples/routing/nethttp_compat/)                        | net/http compatibility: mount handlers and params  |
| [onerror_notfound](examples/platform/onerror_notfound/)                   | Override OnError and NotFound handlers             |
| [otel](examples/platform/otel/)                                           | OpenTelemetry tracing middleware (stdout exporter) |
| [query_params](examples/routing/query_params/)                            | Read query params and respond with JSON            |
| [sessions](examples/middleware/sessions/)                                 | Cookie and header-based sessions                   |
| [static](examples/io/static/)                                             | Serve static files from multiple directories       |
| [templates](examples/io/templates/)                                       | HTML templates rendering                           |
| [validation](examples/validation/validation/)                             | Validate input and return field errors             |
| [validation_with_i18n](examples/validation/validation_with_i18n/)         | Validation with i18n translations                  |
| [websocket](examples/io/websocket/)                                       | WebSocket upgrade and echo                         |

## 🧪 Verify locally

From the repository root:

```bash
# Build all packages in the root module (most examples)
go build ./...

# Build nested module examples (if any)
find examples -type f -name go.mod -execdir sh -c 'go mod tidy && go build ./...' \;
```

## 🛠️ Troubleshooting

- If you see missing modules, run `go mod tidy` in the example folder.
- Some examples depend on third-party libs (e.g., `gorilla/websocket`, `zap`, OpenTelemetry). Go will fetch them automatically.
- Port already in use? Change the port in the example's `main.go`.

Enjoy exploring! ✨
