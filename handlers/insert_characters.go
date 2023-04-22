package handlers

import (
	"encoding/json"
	"fmt"
	"highscore-db/models"
	"log"
	"net/http"
)

func CreateCharacters(w http.ResponseWriter, r *http.Request) {

	var character models.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = models.InsertFromJSON("Taseif_characters.json")

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": "Character inserido com sucesso!",
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
