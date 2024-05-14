package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRandomMeal() (Meal, error) {
	const api = "https://themealdb.com/api/json/v1/1/random.php"

	var meal Meal
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
	if err := json.NewDecoder(resp.Body).Decode(&meals); err != nil {
		return meal, err
	}

	if len(meals.Meals) == 0 {
		return meal, fmt.Errorf("no meals found")
	}

	return meals.Meals[0], nil
}

func GetMealById(id string) (Meal, error) {
	const api = "https://www.themealdb.com/api/json/v1/1/lookup.php?i="

	var meal Meal
	resp, err := http.Get(api + id)
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
	if err := json.NewDecoder(resp.Body).Decode(&meals); err != nil {
		return meal, err
	}

	if len(meals.Meals) == 0 {
		return meal, fmt.Errorf("no meals found")
	}

	return meals.Meals[0], nil
}

func SearchMealsByName(name string) ([]Meal, error) {
	const api = "https://www.themealdb.com/api/json/v1/1/search.php?s="

	resp, err := http.Get(api + name)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var mealsResp mealsResponse
	if err := json.NewDecoder(resp.Body).Decode(&mealsResp); err != nil {
		return nil, err
	}

	if len(mealsResp.Meals) == 0 {
		return nil, fmt.Errorf("no meals found")
	}

	return mealsResp.Meals, nil
}
