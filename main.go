package main

import (
	"fmt"
	"highscore-db/configs"
	"highscore-db/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/", handlers.CreateCharacters)
	r.Post("/1", handlers.Create)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
