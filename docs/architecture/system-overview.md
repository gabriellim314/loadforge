# Load Test Architecture Overview

## Objetivo
LoadTest is a distributed platform for running HTTP load tests.
The system allows users to create tests, monitor their execution in real time, and analyze performance metrics.

## Components
### Dashboard
User interface. Responsible for:
- create tests
- monitor execution
- view metrics

### API
- receive requests from the dashboard
- validate tests
- create jobs
- start workers
- store results

### Worker
Executes the test.
Creates virtual users.
Each virtual user sends HTTP requests.

### Target API
Application used as the test target.
It exists solely for development purposes.



