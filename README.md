# LoadForge

**A lightweight load testing platform built from scratch in Go.**

LoadForge simulates HTTP traffic, measures performance, and helps you understand how an API behaves under load. The project is intentionally built in stages — starting with a solid single-worker engine before any distributed architecture.

---

## Current status — Stage 1 (MVP)

Focus: **one reliable load engine**, not Kubernetes, queues, or multiple workers yet.

```
Frontend          →  🔴 not started
Backend (API)     →  🔴 not started
Load Engine       →  🟡 working (CLI runner + worker)
Target API        →  ✅ minimal /health server
```

### Working today

* ✅ Target API (`GET /health` on `:8000`)
* ✅ Concurrent load generation (goroutine pool + concurrency limit)
* ✅ Per-request timeout via `context.Context`
* ✅ Shared `http.Client`
* ✅ Thread-safe metrics (`sync.Mutex`)
* ✅ Metrics:
  * total / successful / failed requests
  * success rate
  * average latency
  * status code distribution
  * test duration (wall clock)
  * RPS
* ✅ Percentiles (p50 / p95 / p99)
* ✅ CLI flags (URL, count, concurrency still hardcoded in the runner)

### Not yet

* 🔴 Duration-based tests (only total request count)
* 🔴 Backend control API
* 🔴 Result persistence / database
* 🔴 Frontend dashboard
* ⚪ Docker / Kubernetes / RabbitMQ (planned for later stages)

---

## Architecture (today)

```
loadforge/
├── app/
│   ├── target-api/          # sample HTTP target
│   │   └── cmd/server
│   └── load-engine/         # load generator
│       ├── cmd/runner       # CLI entrypoint
│       └── internal/
│           ├── httpclient/  # HTTP + context timeout
│           ├── metrics/     # aggregation + report
│           └── worker/      # concurrency-limited execution
├── docs/architecture/
└── go.mod
```

```
runner
  └── worker (concurrency N)
        └── httpclient  ──HTTP──►  target-api :8000
              │
              └── metrics.Add(...) → Report() on stdout
```

There is **no** control API, queue, or frontend yet. Communication is plain HTTP between the engine and the target.

---

## Getting started

### Requirements

* Go 1.22+

### Clone

```bash
git clone https://github.com/gabriellim314/loadforge.git
cd loadforge
```

### 1. Start the Target API

```bash
go run ./app/target-api/cmd/server
```

Expected:

```
🚀 Target API running on :8000
```

Smoke test:

```bash
curl http://localhost:8000/health
```

### 2. Run the load engine

In another terminal:

```bash
go run ./app/load-engine/cmd/runner
```

Default config (hardcoded in `cmd/runner` for now):

* URL: `http://localhost:8000/health`
* Total requests: `100`
* Concurrency: `10`
* Per-request timeout: `5s`

Example output:

```
Total Requests: 100
Successes: 100
Failures: 0
Average Latency: 1.24ms
Success Rate: 100.00%
Status Code Stats: map[200:100]
RPS: 7545.15
Duration: 13.25ms
```

Optional race check while developing:

```bash
go run -race ./app/load-engine/cmd/runner
```

---

## Roadmap

Aligned with the project stages:

### Stage 1 — MVP (in progress)

* [x] Target API
* [x] HTTP client with latency
* [x] Concurrent worker (bounded concurrency)
* [x] Context + per-request timeout
* [x] Thread-safe metrics (totals, status codes, duration, RPS)
* [x] Percentiles (p50 / p95 / p99)
* [x] Configurable run (CLI flags: URL, count, concurrency, timeout)
* [ ] Thin backend API to start tests and return results
* [ ] Minimal frontend to create a test and view results

### Stage 2 — Multiple workers

Distribute one test across several workers; clear split between control, execution, and metrics aggregation.

### Stage 3 — Docker

Containerize backend, worker, and target API so workers can run independently.

### Stage 4 — Kubernetes

Dynamic worker lifecycle, resource limits, horizontal scale — only after multi-worker + Docker are real.

### Stage 5 — Queue (e.g. RabbitMQ)

Decouple backend from workers for job distribution and failure handling at larger scale.

---

## Design principles

* Correctness and stability before raw performance
* No unbounded goroutine-per-request without a concurrency limit
* Prefer simple, testable Go stdlib solutions
* Do not introduce Docker / K8s / queues before Stage 1 is solid

---

## License

MIT License

## Author

**Gabriel Lima**

Software Engineer | Electrical Engineering Student
