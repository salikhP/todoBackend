package main

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"todoApp/internal/database"
	"todoApp/internal/env"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	dsn := env.GetEnvString("POSTGRES_DSN", "")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-token-2233"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
