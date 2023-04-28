package handlers

import (
	"encoding/json"
	processadados "highscore-db/processa_dados"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func RotinaExec(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Erro ao fazer o parse do id:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = processadados.ExecutaRotina(id)
	if err != nil {
		log.Println("Erro ao executar rotina 1:", err)
	}
	resp := map[string]any{
		"success": true,
		"data":    "Rotina 1 executada com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
