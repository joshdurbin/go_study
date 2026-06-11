# net/http

Go's stdlib HTTP server and client. Production-grade out of the box — no framework needed for most services.

## Server side

- `http.HandleFunc(pattern, fn)` registers a handler on the default mux. For real services, build your own `http.ServeMux` (or use a third-party router) — the default mux is package-global state.
- A handler is `func(w http.ResponseWriter, r *http.Request)`. Write headers before the body; the first `Write` calls `WriteHeader(200)` implicitly.
- `http.ListenAndServe(":8080", mux)` blocks forever. In tests, use `httptest.NewServer(handler)` — it picks a random port, returns a `*Server` with a `URL` field, and shuts down via `Close()`.

## Client side

- `http.Get(url)` is the convenience form. For anything beyond GET, build a request: `http.NewRequest(method, url, body)` then `client.Do(req)`.
- **Always close the response body** (`defer resp.Body.Close()`). Failing to do so leaks file descriptors and keeps the underlying TCP connection out of the pool.
- `http.DefaultClient` has no timeout — fine for demos, dangerous in production. Build a `&http.Client{Timeout: 5 * time.Second}` for real code.

## Worth knowing

- `r.Context()` returns a context tied to the request lifecycle. Cancel-aware downstream calls get cleanup for free when the client disconnects.
- `httptest.NewRecorder()` lets you call a handler directly in tests without a network — fast and deterministic.

## Common gotcha

Forgetting to drain or close the response body. Even on a 4xx/5xx you must `io.Copy(io.Discard, resp.Body); resp.Body.Close()` — otherwise the connection can't be reused, and under load the client starts opening new connections for every request.
