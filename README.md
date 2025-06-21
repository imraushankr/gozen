# GoZen - Complete Backend Framework

## Project Structure
```
gozen/
├── config/
│   └── config.yaml
├── src/
│   ├── api/
│   │   └── routes.go
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── controller/
│   │   │   ├── auth_controller.go
│   │   │   ├── user_controller.go
│   │   │   └── todo_controller.go
│   │   ├── db/
│   │   │   └── db.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   └── rate_limit.go
│   │   ├── models/
│   │   │   ├── user_model.go
│   │   │   ├── todo_model.go
│   │   │   └── base_model.go
│   │   ├── repo/
│   │   │   ├── interfaces.go
│   │   │   ├── user_repo.go
│   │   │   └── todo_repo.go
│   │   ├── routes/
│   │   │   ├── auth_routes.go
│   │   │   ├── user_routes.go
│   │   │   └── todo_routes.go
│   │   ├── security/
│   │   │   ├── hash.go
│   │   │   ├── jwt.go
│   │   │   ├── password_reset.go
│   │   │   └── validator.go
│   │   └── service/
│   │       ├── interfaces.go
│   │       ├── auth_service.go
│   │       ├── user_service.go
│   │       ├── todo_service.go
│   │       └── email_service.go
│   └── pkg/
│       ├── logger/
│       │   └── logger.go
│       ├── response/
│       │   └── response.go
│       └── utils/
│           ├── pagination.go
│           └── helpers.go
├── .env
├── go.mod
└── go.sum
```

**`gozen`** – *The Minimalist Go Backend Framework*

**Description:**
**Gozen** is a clean, minimalist, and modular backend architecture template for Go (Golang) applications. Inspired by principles of simplicity and clarity, it helps developers quickly scaffold production-ready APIs or services with structured layers like `controller`, `service`, and `repo`. Gozen promotes best practices like separation of concerns, configuration via `.env` or YAML, dependency injection, and plug-and-play support for databases and middlewares.

**Key Features:**

* 📁 Structured project layout (`cmd`, `internal`, `pkg`, `config`)
* ⚙️ Supports SQLite, PostgreSQL, MySQL, MongoDB
* 🧩 Easily extendable with REST, gRPC, or GraphQL
* 🛡️ Built-in middleware support (auth, logging, recovery)
* 🚀 Perfect for Clean Architecture or microservice scaffolding

> Build fast. Keep it Zen. ✨

Let me know if you want a GitHub `README.md` intro or folder structure for `gozen`.