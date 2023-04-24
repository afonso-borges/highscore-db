package models

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetNewData(directory, filename string) error {
	cmd := exec.Command("python3", filename)
	cmd.Dir = directory

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao executar o arquivo Python: %v. Mensagem de erro: %s", err, stderr.String())
	}

	return nil
}
