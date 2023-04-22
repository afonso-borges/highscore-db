package models

import (
	"highscore-db/db"
)

func AddExp(character Character) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE characters SET exp=$2 WHERE name=$1`, character.CharacterName, character.Exp)
	if err != nil {
		return 0, nil
	}

	return res.RowsAffected()
}
