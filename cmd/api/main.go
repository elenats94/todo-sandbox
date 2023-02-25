package main

import (
	"log"
	"os"
	"todo-sandbox/internal/app"
)

func main() {
	addr := os.Getenv("APP_ADDR")
	if len(addr) == 0 {
		log.Fatal("$SERVER_ADDR is not set")
	}

	log.Fatal(app.NewApp().Run(addr))
}
