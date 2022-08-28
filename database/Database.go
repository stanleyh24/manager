package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const dbSource = "postgres://root:asd24690@localhost:5432/clientmanager?sslmode=disable"

func ConnectDB() *pgxpool.Pool {

	//db, err := pgx.Connect(context.Background(), dbSource)
	db, err := pgxpool.Connect(context.Background(), dbSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
