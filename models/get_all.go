package models

import "highscore-db/db"

func GetAll() (characters []Character, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM characters ORDER BY dailyexp`)
	if err != nil {
		return
	}

	for rows.Next() {
		var character Character

		err = rows.Scan(&character.ID, &character.CharacterName, &character.Level, &character.GuildIn, &character.Exp, &character.DailyExp)
		if err != nil {
			continue
		}
		if character.DailyExp != 0 {
			characters = append(characters, character)
		}

	}
	return
}
