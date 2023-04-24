package main

import (
	"fmt"
	"highscore-db/configs"
	"highscore-db/handlers"
	processadados "highscore-db/processa_dados"
	"net/http"
	"time"

	"github.com/go-chi/chi"
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
	r.Post("/update_exp", handlers.UpdateAllExp)
	r.Get("/get_characters_list", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
