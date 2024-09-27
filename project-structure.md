event-management-backend/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── middleware.go
│   │   └── service.go
│   ├── db/
│   │   ├── postgres.go
│   │   └── migrate.go
│   ├── event/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── models/
│   │   ├── event.go
│   │   └── user.go
│   ├── router/
│   │   └── router.go
│   ├── user/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   └── repository/
│       ├── base_repository.go
│       └── repository.go
├── pkg/
│   ├── jwt/
│   │   └── jwt.go
│   └── utils/
│       └── utils.go
├── migrations/
│   └── initial_schema.sql
├── scripts/
│   └── migrate.sh
├── .env
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
