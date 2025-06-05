# URL Shortener API - Go + Fiber + Azure Table Storage

This is a simple URL shortener API built in Go using Fiber web framework and Azure Table Storage as the persistence layer.

---

## Purpose

This project is part of my learning journey to:

- Get hands-on with Go (Golang)
- Practice Clean Architecture concepts (still a work in progress)
- Use Fiber for lightweight HTTP routing
- Integrate with Azure Table Storage
- Learn to build backend services with cloud SDKs

---

## Setup

### Prerequisites

- Go 1.20+
- Azure Storage Account with Table enabled
- `.env` file with:

```env
AZURE_STORAGE_ACCOUNT_NAME=your_account_name
AZURE_STORAGE_ACCOUNT_KEY=your_account_key
AZURE_TABLE_NAME=your_table_name

PORT=3000
BASE_URL=http://localhost:3000/
```

### Run Locally
```bash
go mod tidy
go run cmd/server/main.go
```

API will be available at http://localhost:3000/