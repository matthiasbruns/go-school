package file

import (
	"fmt"
	"goschool/0026-recipe-cli-exercise/result/api"
	"io"
	"os"
)

// SaveToFile saves the meal to a file
func formatMeal(meal api.Meal) string {
	return fmt.Sprintf("ID: %s\nName: %s\nCategory: %s\nInstructions: %s\n",
		meal.IDMeal, meal.StrMeal, meal.StrCategory, meal.StrInstructions)
}

func SaveToFile(filename string, meal api.Meal) error {
	// TODO this implementation is ugly, fix it

	fileContent := formatMeal(meal)

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
