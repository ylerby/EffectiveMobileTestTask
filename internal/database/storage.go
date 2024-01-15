package sql

import (
	"EffectiveMobileTask/internal/models"
	"EffectiveMobileTask/schemas"
	"gorm.io/gorm"
)

type Sql struct {
	DB *gorm.DB
}

type InterfaceSql interface {
	Connect() error
	CreateRecord(data *schemas.CreateRecordStruct) error
	ReadRecordByFullName(name, surname, patronymic string) ([]models.RecordModel, error)
	ReadRecordByAge(age int) ([]models.RecordModel, error)
	ReadRecordByGender(gender string) ([]models.RecordModel, error)
	ReadRecordByCountry(country string) ([]models.RecordModel, error)
	UpdateRecordByName(data *schemas.UpdateRecordByNameRequestBody) error
	UpdateRecordBySurname(data *schemas.UpdateRecordBySurnameRequestBody) error
	UpdateRecordByPatronymic(data *schemas.UpdateRecordByPatronymicRequestBody) error
	UpdateRecordByAge(data *schemas.UpdateRecordByAgeRequestBody) error
	DeleteRecord(ID int) error
	GetAllRecords() []models.RecordModel
}
