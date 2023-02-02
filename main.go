package main

import (
	"context"
	"fmt"
	"log"

	//"os"
	"time"

	"jwt-authentication/database"
	"jwt-authentication/test"

	"github.com/jackc/pgx/v4/pgxpool"
	// "github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("postgres://db:secret@localhost:5432/authentication"))
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://db:secret@localhost:5432/authentication")
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	for {
		if err := dbpool.Ping(context.Background()); err != nil {
			fmt.Printf("got error pinging pool, trying again in 10ms\n")
			log.Fatal("error = ", err)
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Println("PING SUCCESS")
		break
	}

	dbRepo := database.NewPostgresRepo(dbpool)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	test.PostgresDemo(ctx, dbRepo)
}
