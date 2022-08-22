package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:asd24690@localhost:5432/clientmanager?sslmode=disable"
	// urlExample:="postgres://username:password@localhost:5432/database_name"
)

func ConnectDB2() *pgx.Conn {

	db, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
