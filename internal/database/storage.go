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
	UpdateRecordByName(data *schemas.UpdateRecordByNameRequestBody) error
	UpdateRecordBySurname(data *schemas.UpdateRecordBySurnameRequestBody) error
	UpdateRecordByAge(data *schemas.UpdateRecordByAgeRequestBody) error
	DeleteRecord(ID int) error
	GetAllRecords() []models.RecordModel
}
