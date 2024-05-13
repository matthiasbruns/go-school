package api

import (
	"fmt"
	"io"
	"net/http"
)

type mealsResponse struct {
	Meals []Meal `json:"meals"`
}

func GetRandomMeal() (Meal, error) {
	var meal Meal
	const api = "https://themealdb.com/api/json/v1/1/random.php"
	resp, err := http.Get(api)
	if err != nil {
		return meal, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return meal, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var meals mealsResponse
	// TODO parse json into &meal

	return meal, nil
}
