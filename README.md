# 🚀 Go Task Master API

A robust, production-ready Task Management RESTful API built with **Golang (Gin)** and **PostgreSQL**. Designed with Clean Architecture principles and containerized with **Docker**.

## 🛠 Tech Stack

* **Language:** Golang (1.23+)
* **Framework:** Gin Gonic
* **Database:** PostgreSQL 15
* **ORM:** GORM
* **Containerization:** Docker & Docker Compose
* **Environment:** Godotenv

## ✨ Features

* **RESTful Endpoints:** Full CRUD operations for task management.
* **Clean Architecture:** Modular structure (Controllers, Models, Routes).
* **Dockerized:** Ready to deploy with a single command.
* **Auto-Migration:** Database tables are automatically generated via GORM.

## 🚀 How to Run

### Option 1: Using Docker (Recommended)
You don't need to install PostgreSQL locally. Just run:

```bash
# Clone the repository
git clone https://github.com/aliemreipek/go-task-master.git
cd go-task-master

# Start the services
docker-compose up -d
```

The API will be available at: `http://localhost:8080`

### Option 2: Manual Run
1.  Ensure PostgreSQL is running locally.
2.  Update `.env` file with your local DB credentials.
3.  Run the application:
    ```bash
    go run cmd/main.go
    ```

## 🔌 API Endpoints

| Method | Endpoint          | Description            |
| :---   | :---              | :---                   |
| `POST` | `/api/v1/tasks`   | Create a new task      |
| `GET`  | `/api/v1/tasks`   | Get all tasks          |
| `GET`  | `/api/v1/tasks/:id`| Get specific task     |
| `PUT`  | `/api/v1/tasks/:id`| Update a task         |
| `DELETE`| `/api/v1/tasks/:id`| Delete a task         |

## 🧪 Testing

You can import the collection to **Postman** or use the included `test_api.http` file if you are using JetBrains IDEs.

```json
// Sample JSON for POST /tasks
{
  "title": "Complete Golang Project",
  "description": "Deploy to GitHub",
  "status": "pending"
}
```

---
*Developed by [Ali Emre Ipek](https://github.com/aliemreipek)*