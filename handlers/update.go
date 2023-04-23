package handlers

import (
	"encoding/json"
	"fmt"
	"highscore-db/models"
	"log"
	"net/http"
)

func UpdateExp(w http.ResponseWriter, r *http.Request) {

	var character models.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{"Taseif_characters.json", "Counterplay_characters.json"}
	var resp map[string]any

	for _, file := range files {
		err := models.UpdateExpFromJSON(file)
		if err != nil {
			resp = map[string]any{
				"Error":   true,
				"Message": fmt.Sprintf("Ocorreu um erro ao tentar atualizar exp: %v", err),
			}
		} else {
			resp = map[string]any{
				"Error":   false,
				"Message": "Exp atualizada com sucesso!",
			}
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UpdateDailyExp(w http.ResponseWriter, r *http.Request) {

	var character models.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{"Taseif_characters.json", "Counterplay_characters.json"}
	var resp map[string]any

	for _, file := range files {
		err := models.UpdateDailyExpFromJSON(file)
		if err != nil {
			resp = map[string]any{
				"Error":   true,
				"Message": fmt.Sprintf("Ocorreu um erro ao tentar atualizar dailyexp: %v", err),
			}
		} else {
			resp = map[string]any{
				"Error":   false,
				"Message": "DailyExp atualizada com sucesso!",
			}
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
