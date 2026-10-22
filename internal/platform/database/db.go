package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	// Register the "pgx" driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(dbrl string) (*sql.DB, error) {

	// interface to connect with driver // connection pooling
	db, err := sql.Open("pgx", dbrl)

	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)

	// fail fast , check database is alive
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Ping: %w", err)
	}
	return db, nil

}
