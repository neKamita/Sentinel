# Sentinel: Real-Time Security Analysis and Anomaly Detection Platform

**Sentinel** is a microservices-based, event-driven platform for collecting, processing, and analyzing event streams in real-time to detect suspicious activity, behavioral anomalies, and potential security threats.

## Architecture

The system is built on a modern, event-driven microservices architecture. Key components communicate asynchronously via a message broker.

### Core Components

- **`ingestor-gateway` (Go):** High-performance gateway for ingesting event data.
- **`core-api` (Java & Spring Boot):** Main API for system management, rules, and user administration.
- **`anomaly-detector` (Python):** Core analysis engine using rules and machine learning to detect threats.
- **`notification-service` (Go):** Service for dispatching alerts via various channels.
- **`dashboard-gateway` (Java & Spring Boot):** Real-time data provider for the monitoring dashboard via WebSockets.

### Infrastructure

- **PostgreSQL 17:** Primary data store for configurations, rules, and incidents.
- **RabbitMQ 4.x:** Message broker for asynchronous communication between services.
- **Redis 8.x:** In-memory data store for caching and real-time event streaming.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Run the Infrastructure

To start the core infrastructure components (PostgreSQL, RabbitMQ, Redis), run:

```bash
docker-compose up -d
```

### Build and Run Services

(Instructions to be added for each service)
