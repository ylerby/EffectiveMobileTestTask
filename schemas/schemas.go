package schemas

type CreateRecordRequestBody struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

type UpdateRecordByNameRequestBody struct {
	ID      int    `json:"id"`
	NewName string `json:"new_name"`
}

type UpdateRecordBySurnameRequestBody struct {
	ID         int    `json:"id"`
	NewSurname string `json:"new_surname"`
}

type UpdateRecordByAgeRequestBody struct {
	ID     int `json:"id"`
	NewAge int `json:"new_age"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CorrectResponse struct {
	Data interface{} `json:"data"`
}

type CreateRecordStruct struct {
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Country    string
}

type GetAgeRequestBody struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GetGenderRequestBody struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type GetCountryRequestBody struct {
	Count   int         `json:"count"`
	Name    string      `json:"name"`
	Country []Countries `json:"country"`
}

type Countries struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type DeleteRecordRequestBody struct {
	ID int `json:"id"`
}
