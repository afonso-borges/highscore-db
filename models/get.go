package models

import (
	"highscore-db/db"
)

func GetAll() (characters []Character, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM characters ORDER BY dailyexp DESC`)
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

func GetGuildExp() (guilds []Guild, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	guildNames := []string{"Taseif", "Counterplay"}

	for _, guildname := range guildNames {
		rows, err := conn.Query(`SELECT guild, SUM(dailyexp) FROM characters WHERE guild = $1 GROUP BY guild`, guildname)
		if err != nil {
			break
		}

		for rows.Next() {
			var guild Guild

			err = rows.Scan(&guild.GuildName, &guild.ExpAmount)
			if err != nil {
				continue
			}
			guilds = append(guilds, guild)
		}
	}
	return
}
