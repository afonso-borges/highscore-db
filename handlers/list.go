package handlers

import (
	"encoding/json"
	"highscore-db/models"
	"log"
	"net/http"
	"strconv"
)

type CharacterInfo struct {
	CharacterName string `json:"CharacterName"`
	Level         int    `json:"level"`
	GuildIn       string `json:"guild"`
	DailyExp      string `json:"dailyexp"`
}

type GuildInfo struct {
	Guild string `json:"guild"`
	Exp   string `json:"dailyexp"`
}

func List(w http.ResponseWriter, r *http.Request) {
	characters, err := models.GetAll()
	if err != nil {
		log.Println("Erro ao obter registros:", err)
	}

	var resp []CharacterInfo
	for _, c := range characters {
		formattedDailyExp := formatDailyExp(c.DailyExp)
		characterInfo := CharacterInfo{
			CharacterName: c.CharacterName,
			Level:         c.Level,
			GuildIn:       c.GuildIn,
			DailyExp:      formattedDailyExp,
		}
		resp = append(resp, characterInfo)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetGuildExp(w http.ResponseWriter, r *http.Request) {
	guildExp, err := models.GetGuildExp()
	if err != nil {
		log.Println("Erro ao obter registros:", err)
	}

	var resp []GuildInfo
	for _, g := range guildExp {
		formattedExp := formatDailyExp(g.ExpAmount)
		guild := GuildInfo{
			Guild: g.GuildName,
			Exp:   formattedExp,
		}
		resp = append(resp, guild)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func formatDailyExp(n int64) string {
	// Separa o sinal do número
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	} else {
		sign = "+"
	}

	// Converte o número para string e obtém o tamanho
	ns := strconv.FormatInt(n, 10)
	l := len(ns)

	// Aplica a formatação, adicionando pontos a cada três dígitos
	formatted := ""
	for i, c := range ns {
		formatted += string(c)
		if (l-i-1)%3 == 0 && i != l-1 {
			formatted += "."
		}
	}

	// Concatena o sinal com a string formatada e retorna
	return sign + formatted
}
