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
cd basic
go run .
```

Most examples start a server on <http://localhost:8080> unless stated otherwise.

## 📚 Examples index

| Example                                               | What it shows                                      |
| ----------------------------------------------------- | -------------------------------------------------- |
| [basic](basic/)                                       | Minimal server and one route                       |
| [any_methods](any_methods/)                           | Handle any HTTP method on a route                  |
| [binding_json](binding_json/)                         | JSON request binding and JSON response             |
| [context](context/)                                   | Store/get values from the per-request context      |
| [context_cancel](context_cancel/)                     | Respect context deadlines and client disconnects   |
| [cookies](cookies/)                                   | Setting and reading cookies                        |
| [cors_timeout](cors_timeout/)                         | CORS and per-route timeout middleware              |
| [csrf](csrf/)                                         | CSRF protection middleware and a simple form       |
| [custom_logger_zap](custom_logger_zap/)               | Using zap via a minimal slog adapter + access logs |
| [error_handling](error_handling/)                     | Custom error handling and 404 responses            |
| [graceful_shutdown](graceful_shutdown/)               | Graceful HTTP server shutdown                      |
| [grouping](grouping/)                                 | Route groups and lightweight middleware            |
| [group_nested](group_nested/)                         | Nested groups with layered middleware              |
| [gzip_requestid_ratelimit](gzip_requestid_ratelimit/) | Gzip + request ID + rate limit + buffer middleware |
| [headers_status](headers_status/)                     | Setting headers and custom HTTP status             |
| [http_mount](http_mount/)                             | Mount net/http handlers on flash and mix routing   |
| [logger_from_context](logger_from_context/)           | Structured logging with slog pulled from context   |
| [method_not_allowed](method_not_allowed/)             | Custom 405 Method Not Allowed handler              |
| [middleware](middleware/)                             | Tiny middleware example adding a header            |
| [nethttp_compat](nethttp_compat/)                     | net/http compatibility: mount handlers and params  |
| [onerror_notfound](onerror_notfound/)                 | Override OnError and NotFound handlers             |
| [otel](otel/)                                         | OpenTelemetry tracing middleware (stdout exporter) |
| [query_params](query_params/)                         | Read query params and respond with JSON            |
| [sessions](sessions/)                                 | Cookie and header-based sessions                   |
| [static](static/)                                     | Serve static files from multiple directories       |
| [templates](templates/)                               | HTML templates rendering                           |
| [validation](validation/)                             | Validate input and return field errors             |
| [validation_with_i18n](validation_with_i18n/)         | Validation with i18n translations                  |
| [websocket](websocket/)                               | WebSocket upgrade and echo                         |

## 🧪 Verify locally

From the repository root:

```bash
# Build all packages in the root module (most examples)
go build ./...

# Build nested module examples too
( cd custom_logger_zap && go build ./... )
( cd otel && go build ./... )
( cd websocket && go build ./... )
```

## 🛠️ Troubleshooting

- If you see missing modules, run `go mod tidy` in the example folder.
- Some examples depend on third-party libs (e.g., `gorilla/websocket`, `zap`, OpenTelemetry). Go will fetch them automatically.
- Port already in use? Change the port in the example's `main.go`.

Enjoy exploring! ✨
