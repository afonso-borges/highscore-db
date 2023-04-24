package processadados

import (
	"fmt"
	"highscore-db/models"
	"time"
)

func FullProcessAlchemist() (int, error) {
	start := time.Now()

	err := models.GetNewData("processa_dados/", "processa_dados.py")
	if err != nil {
		fmt.Println(err)
	}

	files := []string{"processa_dados/data/Taseif_characters.json", "processa_dados/data/Counterplay_characters.json"}
	for _, file := range files {
		fmt.Println("Inserindo novos characters do arquivo", file)
		err := models.InsertFromJSON(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Atualizando Daily Exp da guild", file)
		err = models.UpdateDailyExpFromJSON(file)
		if err != nil {
			fmt.Println(err)
		}
	}
	time_elapsed := time.Since(start)
	fmt.Println("Tempo decorrido:", time_elapsed)
	return fmt.Println("Full Process Concluido")
}
