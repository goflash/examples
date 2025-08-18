# 🚀 GoFlash Examples

[![CI](https://github.com/goflash/examples/actions/workflows/ci.yml/badge.svg)](https://github.com/goflash/examples/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A curated collection of small, focused examples showing how to build web apps with [goflash](https://github.com/goflash/flash).
Each example is a tiny program you can run locally to explore a specific feature or pattern.

## 📦 How to run

- Prerequisites: Go 1.22+ installed.
- Clone this repo, then run any example from its folder.

Example:

```bash
cd categories/routing/basic
go run .
```

Most examples start a server on <http://localhost:8080> unless stated otherwise.

## 📚 Examples by category

### Routing

| Example                                                      | What it shows                                     |
| ------------------------------------------------------------ | ------------------------------------------------- |
| [basic](categories/routing/basic/)                           | Minimal server and one route                      |
| [any_methods](categories/routing/any_methods/)               | Handle any HTTP method on a route                 |
| [grouping](categories/routing/grouping/)                     | Route groups and lightweight middleware           |
| [group_nested](categories/routing/group_nested/)             | Nested groups with layered middleware             |
| [headers_status](categories/routing/headers_status/)         | Setting headers and custom HTTP status            |
| [http_mount](categories/routing/http_mount/)                 | Mount net/http handlers on flash and mix routing  |
| [method_not_allowed](categories/routing/method_not_allowed/) | Custom 405 Method Not Allowed handler             |
| [nethttp_compat](categories/routing/nethttp_compat/)         | net/http compatibility: mount handlers and params |
| [query_params](categories/routing/query_params/)             | Read query params and respond with JSON           |

### Middleware

| Example                                                                     | What it shows                                      |
| --------------------------------------------------------------------------- | -------------------------------------------------- |
| [cors_timeout](categories/middleware/cors_timeout/)                         | CORS and per-route timeout middleware              |
| [csrf](categories/middleware/csrf/)                                         | CSRF protection middleware and a simple form       |
| [gzip_requestid_ratelimit](categories/middleware/gzip_requestid_ratelimit/) | Gzip + request ID + rate limit + buffer middleware |
| [logger_from_context](categories/middleware/logger_from_context/)           | Structured logging with slog pulled from context   |
| [middleware](categories/middleware/middleware/)                             | Tiny middleware example adding a header            |
| [sessions](categories/middleware/sessions/)                                 | Cookie and header-based sessions                   |

### Platform

| Example                                                     | What it shows                                      |
| ----------------------------------------------------------- | -------------------------------------------------- |
| [context](categories/platform/context/)                     | Store/get values from the per-request context      |
| [context_cancel](categories/platform/context_cancel/)       | Respect context deadlines and client disconnects   |
| [custom_logger_zap](categories/platform/custom_logger_zap/) | Using zap via a minimal slog adapter + access logs |
| [error_handling](categories/platform/error_handling/)       | Custom error handling and 404 responses            |
| [graceful_shutdown](categories/platform/graceful_shutdown/) | Graceful HTTP server shutdown                      |
| [onerror_notfound](categories/platform/onerror_notfound/)   | Override OnError and NotFound handlers             |
| [otel](categories/platform/otel/)                           | OpenTelemetry tracing middleware (stdout exporter) |

### I/O

| Example                               | What it shows                                |
| ------------------------------------- | -------------------------------------------- |
| [cookies](categories/io/cookies/)     | Setting and reading cookies                  |
| [static](categories/io/static/)       | Serve static files from multiple directories |
| [templates](categories/io/templates/) | HTML templates rendering                     |
| [websocket](categories/io/websocket/) | WebSocket upgrade and echo                   |

### Validation

| Example                                                             | What it shows                          |
| ------------------------------------------------------------------- | -------------------------------------- |
| [binding_json](categories/validation/binding_json/)                 | JSON request binding and JSON response |
| [validation](categories/validation/validation/)                     | Validate input and return field errors |
| [validation_with_i18n](categories/validation/validation_with_i18n/) | Validation with i18n translations      |

## 🧪 Verify locally

From the repository root:

```bash
./scripts/build.sh
```

## 🛠️ Troubleshooting

- If you see missing modules, run `go mod tidy` in the example folder.
- Some examples depend on third-party libs (e.g., `gorilla/websocket`, `zap`, OpenTelemetry). Go will fetch them automatically.
- Port already in use? Change the port in the example's `main.go`.

Enjoy exploring! ✨
