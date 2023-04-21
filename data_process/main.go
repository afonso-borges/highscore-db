package dataprocess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Chama a função get_player_by_guild
	guildName := "ExemploGuilda"
	if err := get_player_by_guild(guildName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("JSON com personagens da guilda %s criado com sucesso.\n", guildName)

	// Chama a função add_exp_to_JSON
	charactersListPath := fmt.Sprintf("%s_characters.json", guildName)
	world := "Beneva"
	index := 0
	if err := add_exp_to_JSON(charactersListPath, world, index); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Experiência adicionada com sucesso.")

	// Exibe o conteúdo do arquivo JSON
	charactersJSON, err := ioutil.ReadFile(charactersListPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var charactersList []map[string]interface{}
	if err := json.Unmarshal(charactersJSON, &charactersList); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Personagens da guilda:")
	for _, character := range charactersList {
		fmt.Printf("- %s, Level %d, Guild %s, Exp %d\n", character["CharacterName"], int(character["Level"].(float64)), character["Guild"], int(character["Exp"].(float64)))
	}
}
