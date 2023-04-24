package main

import (
	"highscore-db/configs"
	processadados "highscore-db/processa_dados"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	_, err = processadados.FullProcessAlchemist()
	if err != nil {
		return
	}

}
