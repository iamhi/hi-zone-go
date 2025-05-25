package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var IsConnected = false

func ConnectDB(
	username string,
	password string,
	address string,
	database string) {
	dsn := "postgres://" + username + ":" + password + "@" + address + database
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		fmt.Printf("Unable to parse database URL: %v", err)
		return
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		fmt.Errorf("Unable to connect to the database: %v", err)
		return
	}

	DB = pool
	IsConnected = true
	fmt.Println("Connected to database")
}
