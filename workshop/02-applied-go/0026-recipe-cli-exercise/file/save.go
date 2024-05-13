package file

import (
	"fmt"
	"goschool/0026-recipe-cli-exercise/api"
	"io"
	"os"
)

func SaveToFile(filename string, meal api.Meal) error {
	// TODO this implementation is ugly, fix it

	fileContent := fmt.Sprintln(meal)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = io.WriteString(f, fileContent)
	if err != nil {
		return err
	}

	return nil
}
