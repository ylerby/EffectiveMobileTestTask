package service

import (
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAge(name string) (*int, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.agify.io/?name=%s", name), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	currentRequestBody := &schemas.GetAgeRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		return nil, err
	}

	return &currentRequestBody.Age, nil
}

func GetGender(name string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.genderize.io/?name=%s", name), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	currentRequestBody := &schemas.GetGenderRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		return "", err
	}

	return currentRequestBody.Gender, nil
}

func GetCountry(name string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.nationalize.io/?name=%s", name), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	currentRequestBody := &schemas.GetCountryRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		return "", err
	}

	return currentRequestBody.Country[0].CountryID, nil
}
