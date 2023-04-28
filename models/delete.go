package models

import (
	"fmt"
	"highscore-db/db"
)

func DeleteCharFromJson(filename string) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	characters, err := readCharactersFromFile(filename)
	if err != nil {
		return err
	}

	// Faz uma Query de todos os characters no banco de dados
	rows, err := conn.Query(`SELECT * FROM characters`)
	if err != nil {
		return fmt.Errorf("error querying characters from DB: %s", err.Error())
	}
	defer rows.Close()

	charactersInDB := make(map[string]bool)

	for rows.Next() {
		var characterName Character
		if err := rows.Scan(&characterName.ID, &characterName.CharacterName, &characterName.Level, &characterName.GuildIn, &characterName.Exp, &characterName.DailyExp); err != nil {
			return fmt.Errorf("error scanning row: %s", err.Error())
		}
		charactersInDB[characterName.CharacterName] = true
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating over rows: %s", err.Error())
	}

	var numDeleted int
	for _, character := range characters {
		if _, ok := charactersInDB[character.CharacterName]; !ok {
			_, err = conn.Exec("DELETE FROM characters WHERE name = $1", character.CharacterName)
			if err != nil {
				return fmt.Errorf("error deleting character from DB: %s", err.Error())
			}
			numDeleted++
		}
	}

	fmt.Printf(">>>Deleted %d characters from DB\n\n", numDeleted)
	return nil
}
