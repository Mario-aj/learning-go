package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mario-aj/social/internal/env"
	"github.com/mario-aj/social/internal/store"
)

func main() {
	cfg := config{
		address: env.GetString("ADDR", ":8081"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
