package models

import "fmt"

type World struct {
	Name      string
	Guilds    []string
	BattleEye string
}

func (w World) String() string {
	return fmt.Sprintf("World: %s", w.Name)
}

type Character struct {
	Guild
	CharacterName string
	FormerNames   []string
	Level         int
	Exp           int
}

type Guild struct {
	World
	GuildName string
	Players   int
}
