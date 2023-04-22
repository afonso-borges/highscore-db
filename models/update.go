package models

import (
	"fmt"
	"highscore-db/db"
)

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
		// Consulta o valor atual de exp no banco de dados para o personagem atual
		var currentExp int
		err := conn.QueryRow("SELECT exp FROM characters WHERE CharacterName = $1", char.CharacterName).Scan(&currentExp)
		if err != nil {
			return fmt.Errorf("error ao consultar o valor atual de exp para o personagem %s: %v", char.CharacterName, err)
		}

		// Subtrai o novo valor de exp do antigo
		dailyExp := char.Exp - currentExp

		// Atualiza o valor de exp no banco de dados para o novo valor
		_, err = conn.Exec("UPDATE characters SET exp = $1 WHERE CharacterName = $2", char.Exp, char.CharacterName)
		if err != nil {
			return fmt.Errorf("error ao atualizar o valor de exp para o personagem %s: %v", char.CharacterName, err)
		}

		// Insere o novo valor de dailyexp no banco de dados
		_, err = conn.Exec("UPDATE characters SET dailyexp = $1 WHERE CharacterName = $2", dailyExp, char.CharacterName)
		if err != nil {
			return fmt.Errorf("error ao atualizar o valor de dailyexp para o personagem %s: %v", char.CharacterName, err)
		}
	}

	return nil
}
