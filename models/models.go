package models

type Character struct {
	ID            int64  `json:"id"`
	CharacterName string `json:"CharacterName"`
	Level         int    `json:"level"`
	Exp           int    `json:"exp"`
	GuildIn       string `json:"guild"`
}
