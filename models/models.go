package models

type World struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Guilds    []string `json:"guild"`
	BattleEye string   `json:"battle eye"`
}

type Character struct {
	ID            int64  `json:"id"`
	CharacterName string `json:"CharacterName"`
	Level         int    `json:"level"`
	Exp           int    `json:"exp"`
	GuildIn       string `json:"guild"`
}

type Guild struct {
	World
	ID        int64  `json:"id"`
	GuildName string `json:"name"`
	Players   int    `json:"players"`
}
