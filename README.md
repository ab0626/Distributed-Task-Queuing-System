# 🚀 Distributed Task Queue System (Like Celery)

A lightweight, high-performance task queue system built with **Go**, **Redis**, **Docker**, **Prometheus**, and **Grafana**. Designed to showcase systems design, parallelism, and observability.

---

## 🔧 Features

- ✅ REST API to submit tasks
- 🧵 Parallel task execution via worker nodes
- ♻️ Automatic retry policy + dead-letter queue (DLQ)
- 📊 Metrics dashboard with Prometheus + Grafana
- 🐳 Fully containerized using Docker Compose

---

## 🛠️ Tech Stack

- **Go** – Concurrency with goroutines
- **Redis** – Message queue backend
- **Docker** – Containerized API and worker
- **Prometheus** – System metrics
- **Grafana** – Real-time dashboards

---

## 🧪 Getting Started

### 🔁 1. Clone the repo

```bash
git clone https://github.com/your-username/distributed-task-queue.git
cd distributed-task-queue

2. Start the system
bash
Copy
docker-compose up --build

📬 3. Submit a task
bash
Copy
curl -X POST localhost:8080/submit \
-H "Content-Type: application/json" \
-d '{"task_type":"email","payload":"hello@domain.com"}' 

📊 Metrics
localhost:9100/metrics → API metrics

localhost:9101/metrics → Worker metrics

localhost:9090 → Prometheus UI

localhost:3000 → Grafana UI (default login: admin/admin)

📁 Directory Overview
Folder/File	Description
api/	Task submission endpoints (REST)
worker/	Worker node code (process, retry, DLQ)
queue/	Redis queue logic
metrics/	Prometheus metric setup
Dockerfile	API container definition
docker-compose.yml	Defines Redis, API, worker, monitoring
prometheus.yml	Prometheus scrape config
main.go	API entrypoint

📌 Example Metrics
tasks_submitted_total

tasks_processed_total

tasks_failed_total

tasks_retried_total

🚀 Future Improvements
 Task prioritization with Redis sorted sets

 Task scheduling (delayed/cron jobs)

 Secure API auth (JWT / API keys)

 Persistent result storage (PostgreSQL)
