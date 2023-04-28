package processadados

import (
	"fmt"
	"highscore-db/models"
	"time"
)

func Rotina1() (int, error) {
	var file_path string

	file_path = "processa_dados/data/arquivo_unificado.json"

	start := time.Now()
	fmt.Println("Gerando dados...")
	err := models.GetNewData("processa_dados/", "processa_dados.py")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Adicionando novos Characters no DB...")
	err = models.AddNewCharactersFromJson(file_path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Deletando Characters fora da Guild...")
	err = models.DeleteCharFromJson(file_path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Atualizando dados dos Characters...")
	err = models.UpdateExpFromJSON(file_path)
	if err != nil {
		fmt.Println(err)
	}

	time_elapsed := time.Since(start)
	fmt.Println("Tempo decorrido:", time_elapsed)
	return fmt.Println("Rotina 1 concluída")
}

func Rotina2() (int, error) {
	var file_path string

	file_path = "processa_dados/data/arquivo_unificado.json"

	start := time.Now()
	fmt.Println("Gerando dados...")
	err := models.GetNewData("processa_dados/", "processa_dados.py")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Adicionando novos Characters no DB...")
	err = models.AddNewCharactersFromJson(file_path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Deletando Characters fora da Guild...")
	err = models.DeleteCharFromJson(file_path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Atualizando dailyexp dos Characters...")
	err = models.UpdateDailyExpFromJSON(file_path)
	if err != nil {
		fmt.Println(err)
	}

	time_elapsed := time.Since(start)
	fmt.Println("Tempo decorrido:", time_elapsed)
	return fmt.Println("Rotina 2 concluída")
}

func ExecutaRotina(rotina int) error {
	switch rotina {
	case 1:
		_, err := Rotina1()
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	case 2:
		_, err := Rotina2()
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	default:
		return fmt.Errorf("rotina inválida")
	}
	return nil
}
