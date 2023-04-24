package main

import (
	"fmt"
	"highscore-db/configs"
	"highscore-db/handlers"
	processadados "highscore-db/processa_dados"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	_, err = processadados.FullProcessAlchemist()
	if err != nil {
		return
	}
	time.Sleep(time.Second * 1)

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Definir aqui os domínios permitidos
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "application/json"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Adicione o cors como um middleware no seu roteador
	r.Use(c.Handler)

	r.Post("/update_exp", handlers.UpdateAllExp)
	r.Get("/get_characters_list", handlers.List)

	log.Printf("Servidor iniciado em http://localhost:%s", configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
