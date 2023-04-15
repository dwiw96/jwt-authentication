package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	"jwt-authentication/database"
	"jwt-authentication/handlers"

	//"jwt-authentication/test"

	"github.com/jackc/pgx/v4/pgxpool"
	// "github.com/jackc/pgx/v4/pgxpool"
	"github.com/gorilla/mux"
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

	//test.PostgresDemo(ctx, dbRepo)
	handlers.DbConn(ctx, dbRepo)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.AddTable).Methods("POST")
	r.HandleFunc("/admin", handlers.DeleteTable).Methods("POST")
	r.HandleFunc("/user", handlers.AddUser).Methods("POST")
	r.HandleFunc("/admin", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user", handlers.Delete).Methods("DELETE")

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
