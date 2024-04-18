### Golang App Template
An app template to write application in go. It can be rest api, gRPC or whatever.

### Prerequisite
- VSCode
- Go (1.22.0 or later)
- make

### Environment setup
Create a `.env` file. See `.env-example` for reference.

Install goose for database migration.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run `go run ./cmd/api` to start the api server.

The server can be accessed in this url: `http://localhost:3000`

### Migration
All the migration files are stored in `internal/migrations` directory.

Create a new migration file.

```bash
make mg-new name=user_table
```

Apply all available migrations

```bash
make mg-up
```

Revert single migration from current version

```bash
make mg-down
```

To check migration status.

```bash
make mg-status
```

_Note: For local database we might need to remove `tls=skip-verify` from the connection string_