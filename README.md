## Build
# Books API
A simple Go REST API built with a clean layered architecture.
This project is designed as a DevOps-friendly backend service for learning and demonstrating testing, Docker, CI/CD, and deployment practices.
## Features
- Go HTTP server
- Environment-based configuration
- Structured logging with zerolog
- Health check endpoint
- Layered architecture:
  - Handler
  - Service
  - Repository
- Unit tests for service and handler layers
## Project Structure
```text
.
├── cmd/api
│   └── main.go
├── internal
│   ├── handler
│   ├── repository
│   └── service
├── pkg/config
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
## Requirements
- Go 1.22+
- Linux/macOS terminal
## Setup
Clone the repository:
```bash
git clone https://github.com/salar-nourani/books-api.git
cd books-api
```
Create your environment file:
```bash
cp .env.example .env
```
Run the application:
```bash
go run ./cmd/api
```
Or using Make:
```bash
make run
```
## Endpoints
**Health check:**
```bash
curl localhost:8080/health
```
Expected output:
```text
OK
```
**Application status:**
```bash
curl localhost:8080/status
```
Expected output:
```text
Database is connected and healthy
```
## Tests
Run all tests:
```bash
go test ./...
```
Or using Make:
```bash
make test
```
## Build
Build the application binary:
```bash
make build
```
The binary will be created under:
```text
bin/books-api
```
## Clean
Remove build artifacts:
```bash
make clean
```
## Roadmap
- Dockerfile
- Docker Compose
- GitHub Actions CI
- PostgreSQL integration
- Kubernetes manifests
- Helm chart
- Monitoring with Prometheus and Grafana
## Purpose
This repository is part of a hands-on DevOps learning path.
The goal is to build a small but professional backend service and gradually add real-world DevOps practices such as testing, containerization, CI/CD, deployment, and monitoring.
### Build the application binary
```bash
make build
```
The binary will be created under:
```text
bin/books-api
```
## Clean
### Remove build artifacts
```bash
make clean
```
## Roadmap
- Dockerfile
- Docker Compose
- GitHub Actions CI
- PostgreSQL integration
- Kubernetes manifests
- Helm chart
- Monitoring with Prometheus and Grafana
## Purpose
This repository is part of a hands-on DevOps learning path.
The goal is to build a small but professional backend service and gradually add real-world DevOps practices such as testing, containerization, CI/CD, deployment, and monitoring.
