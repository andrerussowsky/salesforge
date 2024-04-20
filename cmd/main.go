package main

import (
    "database/sql"
    "log"
    "net/http"

    "salesforge/internal/handlers"
    "salesforge/internal/repository"
    "salesforge/pkg/server"

    _ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
    // Initialize database connection
    db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/salesforge?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    // Initialize repository
    repository.InitDB(db)

    // Initialize server
    srv := server.NewServer()

    // Define API endpoints
    srv.Router.HandleFunc("/sequence", handlers.CreateSequence).Methods("POST")
    srv.Router.HandleFunc("/sequence/{sequenceID}/step/{stepID}", handlers.UpdateSequenceStep).Methods("PUT")
    srv.Router.HandleFunc("/sequence/{sequenceID}/step/{stepID}", handlers.DeleteSequenceStep).Methods("DELETE")
    srv.Router.HandleFunc("/sequence/{sequenceID}/tracking", handlers.UpdateSequenceTracking).Methods("PUT")

    // Start server
    log.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
