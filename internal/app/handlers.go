package app

import (
	"EffectiveMobileTask/internal/service"
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	age, err := service.GetAge(currentRequestBody.Name)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении возраста - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	}

	gender, err := service.GetGender(currentRequestBody.Name)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении пола - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	}

	country, err := service.GetCountry(currentRequestBody.Name)
	if err != nil {
		message = fmt.Sprintf("ошибка при получении национальности - %s", err)
		a.Logger.Info(message)
		a.ErrorResponseWriter(w, http.StatusServiceUnavailable, message)
		return
	}

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
		Age:        *age,
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

func (a *App) ReadRecordHandler(_ http.ResponseWriter, _ *http.Request) {}

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
