package service

import (
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func GetAge(name string, errCh chan error, valueCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.agify.io/?name=%s", name), nil)
	if err != nil {
		errCh <- err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		errCh <- err
	}
	currentRequestBody := &schemas.GetAgeRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		errCh <- err
	}

	valueCh <- currentRequestBody.Age
}

func GetGender(name string, errCh chan error, valueCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.genderize.io/?name=%s", name), nil)
	if err != nil {
		errCh <- err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		errCh <- err
	}

	currentRequestBody := &schemas.GetGenderRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		errCh <- err
	}

	valueCh <- currentRequestBody.Gender
}

func GetCountry(name string, errCh chan error, valueCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.nationalize.io/?name=%s", name), nil)
	if err != nil {
		errCh <- err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		errCh <- err
	}

	currentRequestBody := &schemas.GetCountryRequestBody{}

	err = json.NewDecoder(response.Body).Decode(currentRequestBody)
	if err != nil {
		errCh <- err
	}

	valueCh <- currentRequestBody.Country[0].CountryID
}
