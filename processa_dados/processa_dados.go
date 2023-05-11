package processadados

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GuildMember struct {
	CharacterName string `json:"CharacterName"`
	Level         int    `json:"level"`
	Guild         string
	Exp           int `json:"exp,omitempty"`
}

type HighScoreApiResponse struct {
	HighScores struct {
		HighScoreListAPI []struct {
			Level    int    `json:"level"`
			Name     string `json:"name"`
			Rank     int    `json:"rank"`
			Title    string `json:"title"`
			Value    int    `json:"value"`
			Vocation string `json:"vocation"`
			World    string `json:"world"`
		} `json:"highscore_list"`
	} `json:"highscores"`
}

type GuildApiResponse struct {
	Guild struct {
		Members []struct {
			Joined   string `json:"joined"`
			Level    int    `json:"level"`
			Name     string `json:"name"`
			Rank     string `json:"rank"`
			Status   string `json:"status"`
			Title    string `json:"title"`
			Vocation string `json:"vocation"`
		} `json:"members"`
	} `json:"guild"`
}

func getGuildMembers(guildName string) ([]GuildMember, error) {
	var members []GuildMember

	url := fmt.Sprintf("https://dev.tibiadata.com/v4/guild/%s", guildName)
	resp, err := http.Get(url)
	if err != nil {
		return members, err
	}
	defer resp.Body.Close()

	var data GuildApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return members, err
	}

	for _, member := range data.Guild.Members {
		if member.Level > 556 {
			memberInfo := GuildMember{
				CharacterName: member.Name,
				Level:         member.Level,
				Guild:         guildName,
				Exp:           0,
			}
			members = append(members, memberInfo)
		}
	}

	return members, nil
}

func insertCharacterExp(characters []GuildMember, world string, index int) ([]GuildMember, error) {
	url := fmt.Sprintf("https://dev.tibiadata.com/v4/highscores/%s/experience/all/%d", world, index)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var data1 HighScoreApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data1)
	if err != nil {
		fmt.Print(err)
	}

	for i, char := range characters {
		for _, highscoreChar := range data1.HighScores.HighScoreListAPI {
			if highscoreChar.Name == char.CharacterName {
				characters[i].Exp = highscoreChar.Value
				break
			}
		}
	}

	return characters, nil

}

func createJson(members []GuildMember, guildName string) error {
	path := fmt.Sprintf("./processa_dados/data/%s_characters.json", guildName)

	byteMembers, err := json.Marshal(members)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, byteMembers, 0755)
	if err != nil {
		return err
	}

	return nil
}

func UnifyJson() error {
	taseifFile, err := os.Open("./processa_dados/data/Taseif_characters.json")
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo Taseif_characters.json: %v", err)
	}
	defer taseifFile.Close()

	var taseifMembers []map[string]interface{}
	err = json.NewDecoder(taseifFile).Decode(&taseifMembers)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o arquivo Taseif_characters.json: %v", err)
	}

	counterplayFile, err := os.Open("./processa_dados/data/Counterplay_characters.json")
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo Counterplay_characters.json: %v", err)
	}
	defer counterplayFile.Close()

	var counterplayMembers []map[string]interface{}
	err = json.NewDecoder(counterplayFile).Decode(&counterplayMembers)
	if err != nil {
		return fmt.Errorf("erro ao decodificar o arquivo Counterplay_characters.json: %v", err)
	}

	gather_data := make([]map[string]interface{}, 0)

	gather_data = append(gather_data, taseifMembers...)
	gather_data = append(gather_data, counterplayMembers...)

	gather_dataFile, err := os.Create("./processa_dados/data/arquivo_unificado.json")
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo arquivo_unificado.json: %v", err)
	}
	defer gather_dataFile.Close()

	err = json.NewEncoder(gather_dataFile).Encode(&gather_data)
	if err != nil {
		return fmt.Errorf("erro ao codificar o arquivo arquivo_unificado.json: %v", err)
	}

	return nil

}

func FullProcessAlchemist() error {
	guildsFP := [2]string{"Taseif", "Counterplay"}

	for _, guild := range guildsFP {
		members, err := getGuildMembers(guild)
		if err != nil {
			return err
		}
		var membersWithExp []GuildMember
		for i := 0; i < 21; i++ {
			if i > 0 {
				membersWithExp, err = insertCharacterExp(members, "Inabra", i)
				if err != nil {
					fmt.Print(err)
					break
				}
			}
		}
		err = createJson(membersWithExp, guild)
		if err != nil {
			return err
		}
	}
	err := UnifyJson()
	if err != nil {
		return err
	}

	return nil
}
