package dataprocess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func get_player_by_guild(guild_name string) bool {
	url := fmt.Sprintf("https://dev.tibiadata.com/v4/guild/%s", guild_name)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	data := make(map[string]map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return false
	}

	guild_members := []map[string]interface{}{}
	for _, member := range data["guild"]["members"].([]interface{}) {
		member_name := member.(map[string]interface{})["name"].(string)
		member_level := member.(map[string]interface{})["level"].(float64)

		member_info := map[string]interface{}{
			"CharacterName": member_name,
			"Level":         member_level,
			"Guild":         guild_name,
		}

		guild_members = append(guild_members, member_info)
	}

	_json, err := json.Marshal(guild_members)
	if err != nil {
		fmt.Println(err)
		return false
	}

	file, err := os.Create(fmt.Sprintf("%s_characters.json", guild_name))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	file.Write(_json)

	if _, err := os.Stat(fmt.Sprintf("%s_characters.json", guild_name)); os.IsNotExist(err) {
		return false
	}

	return true
}

func add_exp_to_JSON(characters_list_path string, world string, index int) {
	file, err := os.OpenFile(characters_list_path, os.O_RDWR, 0644)
	if err != nil {
		return false
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}

	var characters_list []map[string]interface{}
	err = json.Unmarshal(byteValue, &characters_list)
	if err != nil {
		return false
	}

	url := "https://dev.tibiadata.com/v4/highscores/" + world + "/experience/all/" + strconv.Itoa(index)
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}

	var data map[string]map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false
	}

	for _, highscore_char := range data["highscores"]["highscore_list"].([]interface{}) {
		for _, _char := range characters_list {
			if _char["CharacterName"] == highscore_char.(map[string]interface{})["name"] {
				_char["Exp"] = highscore_char.(map[string]interface{})["value"]
			}
		}
	}

	output, err := json.Marshal(characters_list)
	if err != nil {
		return false
	}

	err = ioutil.WriteFile(characters_list_path, output, 0644)
	if err != nil {
		return false
	}

	return true
}
