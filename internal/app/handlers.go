package app

import (
	"EffectiveMobileTask/internal/service"
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func (a *App) CreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	var message string
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.CreateRecordRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.Name == "" || currentRequestBody.Surname == "" {
		message = fmt.Sprintf("некорректные значения полей")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	errorCh := make(chan error)
	wg := &sync.WaitGroup{}

	var age int
	ageValueCh := make(chan int)

	wg.Add(1)
	go service.GetAge(currentRequestBody.Name, errorCh, ageValueCh, wg)
	select {
	case err = <-errorCh:
		message = fmt.Sprintf("ошибка при получении возраста - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	case age = <-ageValueCh:
		message = fmt.Sprintf("успешно получен возраст - %d", age)
		a.Logger.Info(message)
	}

	var gender string
	genderValueCh := make(chan string)

	wg.Add(1)
	go service.GetGender(currentRequestBody.Name, errorCh, genderValueCh, wg)
	select {
	case err = <-errorCh:
		message = fmt.Sprintf("ошибка при получении пола - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	case gender = <-genderValueCh:
		message = fmt.Sprintf("успешно получен пол - %s", gender)
		a.Logger.Info(message)
	}

	var country string
	countryValueCh := make(chan string)

	wg.Add(1)
	go service.GetCountry(currentRequestBody.Name, errorCh, countryValueCh, wg)
	select {
	case err = <-errorCh:
		message = fmt.Sprintf("ошибка при получении национальности - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	case country = <-countryValueCh:
		message = fmt.Sprintf("успешно получена национальность - %s", country)
		a.Logger.Info(message)
	}

	wg.Wait()

	var currentPatronymic string
	if currentRequestBody.Patronymic != nil {
		currentPatronymic = *currentRequestBody.Patronymic
	} else {
		currentPatronymic = ""
	}

	data := &schemas.CreateRecordStruct{
		Name:       currentRequestBody.Name,
		Surname:    currentRequestBody.Surname,
		Patronymic: currentPatronymic,
		Age:        age,
		Gender:     gender,
		Country:    country,
	}

	err = a.Sql.CreateRecord(data)
	if err != nil {
		message = fmt.Sprintf("ошибка при создании записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно создана запись с данными: %v", data)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно создана")
}

func (a *App) ReadRecordByFullNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	patronymic := r.URL.Query().Get("patronymic")

	var message string

	if name == "" || surname == "" {
		message = fmt.Sprintf("некорректное значение параметра(ов)")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	result, err := a.Sql.ReadRecordByFullName(name, surname, patronymic)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении записей")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if len(result) == 0 {
		a.Logger.Debugf("data: %v", result)
		message = "Нет значений, удовлетворяющих условию"
		a.CorrectResponseWriter(w, http.StatusOK, message)
		return
	}

	a.Logger.Debugf("data: %v", result)
	a.CorrectResponseWriter(w, http.StatusOK, result)
}

func (a *App) ReadRecordByAgeHandler(w http.ResponseWriter, r *http.Request) {
	age := r.URL.Query().Get("age")

	var message string

	ageNumber, err := strconv.Atoi(age)
	if err != nil || ageNumber < 0 {
		message = fmt.Sprintf("некорректное значение возраста")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	result, err := a.Sql.ReadRecordByAge(ageNumber)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении записей")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if len(result) == 0 {
		a.Logger.Debugf("data: %v", result)
		message = "Нет значений, удовлетворяющих условию"
		a.CorrectResponseWriter(w, http.StatusOK, message)
		return
	}

	a.Logger.Debugf("data: %v", result)
	a.CorrectResponseWriter(w, http.StatusOK, result)
}

func (a *App) ReadRecordByGenderHandler(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gender")

	var message string

	if gender == "" {
		message = fmt.Sprintf("некорректное значение пола")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	result, err := a.Sql.ReadRecordByGender(gender)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении записей")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if len(result) == 0 {
		a.Logger.Debugf("data: %v", result)
		message = "Нет значений, удовлетворяющих условию"
		a.CorrectResponseWriter(w, http.StatusOK, message)
		return
	}

	a.Logger.Debugf("data: %v", result)
	a.CorrectResponseWriter(w, http.StatusOK, result)
}

func (a *App) ReadRecordByCountryHandler(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")

	var message string

	if country == "" {
		message = fmt.Sprintf("некорректное значение национальности")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	result, err := a.Sql.ReadRecordByCountry(country)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении записей")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if len(result) == 0 {
		a.Logger.Debugf("data: %v", result)
		message = "Нет значений, удовлетворяющих условию"
		a.CorrectResponseWriter(w, http.StatusOK, message)
		return
	}

	a.Logger.Debugf("data: %v", result)
	a.CorrectResponseWriter(w, http.StatusOK, result)
}

func (a *App) UpdateRecordByNameHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.UpdateRecordByNameRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.NewName == "" {
		message = fmt.Sprintf("некорректное значение поля")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	err = a.Sql.UpdateRecordByName(currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при обновлении записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно обновлена запись с данными: %v", currentRequestBody)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно обновлена")
}

func (a *App) UpdateRecordBySurnameHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.UpdateRecordBySurnameRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.NewSurname == "" {
		message = fmt.Sprintf("некорректное значение поля")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	err = a.Sql.UpdateRecordBySurname(currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при обновлении записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно обновлена запись с данными: %v", currentRequestBody)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно обновлена")
}

func (a *App) UpdateRecordByPatronymicHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.UpdateRecordByPatronymicRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.NewPatronymic == "" {
		message = fmt.Sprintf("некорректное значение поля")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	err = a.Sql.UpdateRecordByPatronymic(currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при обновлении записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно обновлена запись с данными: %v", currentRequestBody)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно обновлена")
}

func (a *App) UpdateRecordByAgeHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.UpdateRecordByAgeRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.NewAge < 0 {
		message = fmt.Sprintf("некорректное значение поля")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	err = a.Sql.UpdateRecordByAge(currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при обновлении записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно обновлена запись с данными: %v", currentRequestBody)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно обновлена")
}

func (a *App) DeleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	var message string
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		message = fmt.Sprintf("ошибка чтения тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	currentRequestBody := &schemas.DeleteRecordRequestBody{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		message = fmt.Sprintf("ошибка при десериализации тела запроса - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	if currentRequestBody.ID < 0 {
		message = fmt.Sprintf("некорректное значение поля")
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	err = a.Sql.DeleteRecord(currentRequestBody.ID)
	if err != nil {
		message = fmt.Sprintf("ошибка при удалении записи - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusInternalServerError, message)
		return
	}

	message = fmt.Sprintf("Успешно удалена запись с ID: %d", currentRequestBody.ID)
	a.Logger.Infof(message)
	a.CorrectResponseWriter(w, http.StatusOK, "Запись успешно удалена")
}

func (a *App) GetAllRecordsHandler(w http.ResponseWriter, _ *http.Request) {
	var message string
	result := a.Sql.GetAllRecords()

	if len(result) == 0 {
		a.Logger.Debugf("data: %v", result)
		message = "Таблица пуста"
		a.CorrectResponseWriter(w, http.StatusOK, message)
		return
	}

	a.Logger.Debugf("data: %v", result)
	a.CorrectResponseWriter(w, http.StatusOK, result)
}
