package handlers

import (
	"encoding/json"
	"highscore-db/models"
	"log"
	"net/http"
)

func UpdateAllExp(w http.ResponseWriter, r *http.Request) {
	var character models.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		log.Println("Erro ao fazer decode do json:", err)
	}

	err = models.UpdateExpFromJSON("processa_dados/data/Taseif_characters.json")
	if err != nil {
		log.Println("Erro atualizar exp:", err)
	}
	err = models.UpdateExpFromJSON("processa_dados/data/Counterplay_characters.json")
	if err != nil {
		log.Println("Erro atualizar exp:", err)
	}
	log.Println("Exp atualizada")

	resp := map[string]any{
		"success": true,
		"data":    nil,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
