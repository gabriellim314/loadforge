# LoadForge

**A lightweight, developer-friendly load testing platform built from scratch in Go.**

LoadForge is an open-source load testing tool designed to simulate traffic, measure system performance, and provide meaningful insights about application behavior under different workloads.

The goal of LoadForge is to provide a simple yet powerful alternative for developers who want to understand how their systems behave under pressure, while exploring concepts such as concurrency, distributed systems, networking, and performance engineering.

---

## 🚀 Features

### Current

* ✅ HTTP load generation
* ✅ Target API simulation environment
* ✅ Request latency measurement
* ✅ Success and failure tracking
* ✅ Concurrent request execution using Go routines

### In Development

* 🔄 Virtual users simulation
* 🔄 Requests per second (RPS) control
* 🔄 Configurable load scenarios
* 🔄 Performance reports
* 🔄 Real-time monitoring dashboard
* 🔄 Distributed load generation

---

# 🏗️ Architecture

LoadForge is divided into two main components:

```
loadforge/
│
├── app/
│   │
│   ├── target-api/
│   │   └── A sample HTTP service used as a load testing target
│   │
│   └── load-engine/
│       └── The engine responsible for generating traffic
│
└── go.mod
```

## Target API

The target API is a simple Go HTTP server that provides endpoints for testing.

Example:

```
GET /health
```

Response:

```
200 OK

OK
```

This allows LoadForge to test real HTTP communication patterns.

---

## Load Engine

The Load Engine is the core of LoadForge.

It is responsible for:

* Creating HTTP requests
* Controlling traffic generation
* Measuring response times
* Collecting performance metrics

Example workflow:

```
                 HTTP Requests

+-------------+ ---------------> +-------------+
| LoadForge   |                  | Target API  |
| Engine      | <--------------- |             |
+-------------+    Responses     +-------------+

        |
        |
        v

Performance Metrics

- Requests per second
- Average latency
- Error rate
- Percentiles
```

---

# 🧠 Engineering Concepts

LoadForge is built to explore and implement real-world backend engineering concepts.

## Concurrency

Go routines are used to simulate multiple users accessing a system simultaneously.

Example:

```
Virtual User 1 ───┐
Virtual User 2 ───┤
Virtual User 3 ───┼──> API
Virtual User 4 ───┘
```

This allows realistic workload simulation.

---

## Performance Metrics

LoadForge tracks important performance indicators:

### Latency

The time required for a request to complete.

Example:

```
Request started
      |
      |
Response received

Latency = 15ms
```

---

### Throughput

The amount of traffic processed by the system.

Example:

```
500 requests / second
```

---

### Error Rate

Percentage of failed requests.

Example:

```
Requests: 10000

Successful: 9980
Failed: 20

Error rate: 0.2%
```

---

# 🛠️ Tech Stack

## Backend

* Go
* Standard HTTP library
* Goroutines
* Channels

## Planned

* React dashboard
* WebSocket-based monitoring
* Persistent test history
* Distributed workers

---

# ⚙️ Getting Started

## Requirements

* Go 1.22+

---

## Clone the repository

```bash
git clone https://github.com/gabriellim314/loadforge.git

cd loadforge
```

---

# Running the Target API

Start the testing environment:

```bash
go run ./app/target-api/cmd/server
```

Expected output:

```
🚀 Target API running on :8080
```

Test:

```bash
curl http://localhost:8080/health
```

Response:

```
OK
```

---

# Running LoadForge

Start the load engine:

```bash
go run ./app/load-engine/cmd/runner
```

Example output:

```
Starting load test...

Requests: 1000
Successful: 1000
Failed: 0

Average latency:
12ms
```

---

# 📊 Future Roadmap

## Phase 1 — Core Engine

* [x] HTTP requests
* [x] Latency measurement
* [ ] Concurrent workers
* [x] Metrics collector

## Phase 2 — Advanced Load Testing

* [ ] Virtual users
* [ ] RPS control
* [ ] Custom scenarios
* [ ] Test configurations

## Phase 3 — Platform

* [ ] Web dashboard
* [ ] Historical reports
* [ ] Authentication
* [ ] Cloud execution

## Phase 4 — Distributed Testing

* [ ] Multiple load generators
* [ ] Kubernetes support
* [ ] Large-scale simulations

---

# 🎯 Motivation

Modern applications need to handle unpredictable traffic, and understanding system limits is a fundamental part of building reliable software.

LoadForge was created to study and implement:

* Backend engineering
* Distributed systems
* Network programming
* Performance optimization
* Cloud infrastructure concepts

The project aims to evolve into a complete performance testing platform built with modern engineering principles.

---

# 📄 License

MIT License

---

# 👨‍💻 Author

**Gabriel Lima**

Software Engineer | Electrical Engineering Student

Building systems at the intersection of software, infrastructure, and performance engineering.
