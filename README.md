# GOZEN – The Minimalist Go Backend Framework

![GoZen Logo](https://res.cloudinary.com/cloud-alpha/image/upload/c_pad,b_gen_fill,w_150,h_84,ar_16:9/v1752353216/Common/ChatGPT_Image_Jul_13_2025_02_15_15_AM_zrc03u.png)  
*A lightweight, high-performance backend framework built with Go*

## 🚀 Features

- **Minimalist Design**: Only what you need, nothing more
- **Production Ready**: Built-in configurations for server, database, logging
- **Modular Architecture**: Clean separation of concerns
- **JWT Authentication**: Ready-to-use auth system
- **SQLite First**: Perfect for small to medium applications
- **Simple Migrations**: Easy database schema management

## 📦 Project Structure

```text
gozen/
├── src/
│   ├── cmd/
│   │   ├── migrate/          # Migration commands
│   │   └── server/           # Main application server
│   ├── configs/              # Configuration files
│   │   ├── app.yaml          # Main configuration
│   │   └── config.go         # Config loader
│   └── internal/
│       ├── app/              # Application core
│       ├── delivery/         # Transport layers
│       │   ├── grpc/         # gRPC handlers
│       │   └── http/         # HTTP handlers
│       ├── domain/           # Business entities
│       ├── repository/       # Data access layer
│       └── usecase/          # Business logic
├── migrations/               # Database migrations
├── .env.example              # Environment template
├── go.mod                    # Go dependencies
└── README.md                 # This file
```

## 🛠️ Getting Started

### Prerequisites

- Go 1.20+
- SQLite (or your preferred database)

### Installation

```bash
git clone git@github.com:imraushankr/gozen.git
cd gozen
cp .env.example .env
go mod download
```

### Configuration

Edit `.env` file:

```env
APP_ENV=development
DB_TYPE=sqlite
DB_NAME=gozen.db
```

### Running the Server

```bash
go run src/cmd/server/main.go
```

## 🔧 Migration System

Create new migration:

```bash
go run src/cmd/migrate/create/main.go create_migration_name
```

Run migrations:

```bash
go run src/cmd/migrate/run/main.go
```

## 🛠️ Taskfile Commands

The project includes a Taskfile.yml with useful commands. Install [Task](https://taskfile.dev/) to use them:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

Then you can run these commands:

```bash
# Run the application (with hot-reload)
task run

# Build the application
task build

# Run all tests
task test

# Create new migration
task migration-create -- CLI_ARGS="migration_name"

# Run migrations
task migrate-up

# Rollback last migration
task migrate-down

# Show migration status
task migrate-status

# Tidy Go modules
task tidy

# Run linters
task lint

# Setup development environment
task setup

# Run CI checks
task ci
```

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## ✉️ Contact

Raushan Kumar - [@imraushankr](https://github.com/imraushankr)

---

*"Simplicity is the ultimate sophistication."* - *Leonardo da Vinci*