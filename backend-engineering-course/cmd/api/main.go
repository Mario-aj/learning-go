package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mario-aj/social/internal/env"
)

func main() {
	cfg := config{
		address: env.GetString("ADDR", ":8081"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
