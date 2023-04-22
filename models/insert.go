package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"highscore-db/db"
	"io/ioutil"
	"os"
)

func InsertFromJSON(filename string) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	characters, err := readCharactersFromFile(filename)
	if err != nil {
		return err
	}

	if len(characters) == 0 {
		return fmt.Errorf("o arquivo JSON não contém personagens")
	}

	err = InsertCharacters(characters, conn)
	if err != nil {
		return fmt.Errorf("ocorreu um erro ao inserir personagens: %v", err)
	}

	return nil
}

func readCharactersFromFile(filename string) ([]Character, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var characters []Character
	err = json.Unmarshal(byteValue, &characters)
	if err != nil {
		return nil, err
	}

	return characters, nil
}

func InsertCharacters(characters []Character, db *sql.DB) error {
	for _, char := range characters {
		_, err := db.Exec("INSERT INTO characters (name, guild, level, exp) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING",
			char.CharacterName, char.GuildIn, char.Level, char.Exp)
		if err != nil {
			return fmt.Errorf("error inserting character %s: %s", char.CharacterName, err.Error())
		}
	}
	return nil
}
