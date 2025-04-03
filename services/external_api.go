package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetAge(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Age int `json:"age"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Age, nil
}

func GetGender(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Gender string `json:"gender"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Gender, nil
}

func GetNationality(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Nationality string `json:"country_id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Nationality == "" {
		return "", errors.New("no nationality found")
	}

	return result.Nationality, nil
}
