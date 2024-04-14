package database

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/logging"
)

type Sql struct {
	*sqlx.DB
}

func New(ctx context.Context, cfg *config.Db) (*Sql, error) {
	logger := logging.FromContext(ctx)

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	mysqlCfg := mysql.Config{
		User:                 cfg.Username,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               cfg.Name,
		AllowNativePasswords: true,
		TLSConfig:            "skip-verify",
		ParseTime:            true,
	}

	connectionString := mysqlCfg.FormatDSN()

	logger.Debugln(connectionString)

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		logger.Fatalf("failed to connect to db. error: %w", err)
		return nil, fmt.Errorf("failed to connect to db. error: %w", err)
	}

	logger.Info("connected to database")

	return &Sql{db}, nil
}
