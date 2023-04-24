package handlers

import (
	"encoding/json"
	"highscore-db/models"
	"log"
	"net/http"
)

type CharacterInfo struct {
	CharacterName string `json:"CharacterName"`
	Level         int    `json:"level"`
	GuildIn       string `json:"guild"`
	DailyExp      int64  `json:"dailyexp"`
}

func List(w http.ResponseWriter, r *http.Request) {
	characters, err := models.GetAll()
	if err != nil {
		log.Println("Erro ao obter registros:", err)
	}

	var charactersInfo []CharacterInfo
	for _, c := range characters {
		characterInfo := CharacterInfo{
			CharacterName: c.CharacterName,
			Level:         c.Level,
			GuildIn:       c.GuildIn,
			DailyExp:      c.DailyExp,
		}
		charactersInfo = append(charactersInfo, characterInfo)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charactersInfo)
}
