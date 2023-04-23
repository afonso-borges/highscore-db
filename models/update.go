package models

import (
	"fmt"
	"highscore-db/db"
)

func UpdateDailyExpFromJSON(filename string) error {
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

	for _, char := range characters {
		_, err = conn.Exec("UPDATE characters SET dailyexp=$1 - exp WHERE name=$2", char.Exp, char.CharacterName)
		if err != nil {
			return fmt.Errorf("error ao atualizar o valor de dailyexp para o personagem %s: %v", char.CharacterName, err)
		}
	}

	return nil
}

func UpdateExpFromJSON(filename string) error {
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

	for _, char := range characters {

		// Atualiza o valor de exp no banco de dados para o novo valor
		_, err = conn.Exec("UPDATE characters SET level=$1, exp=$2 WHERE name=$3", char.Level, char.Exp, char.CharacterName)
		if err != nil {
			return fmt.Errorf("error ao atualizar o valor de exp para o personagem %s: %v", char.CharacterName, err)
		}
	}

	return nil
}
