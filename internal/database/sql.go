package sql

import (
	"EffectiveMobileTask/internal/models"
	"EffectiveMobileTask/schemas"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func NewDatabase() *Sql {
	return &Sql{}
}

func (s *Sql) Connect() error {
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("невозможно получить текущую директорию - %s", err)
	}

	envPath := filepath.Join(workingDir, "..", ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("невозможно получить .env-файл - %s", err)
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("ошибка при подключении к БД - %s", err)
	}

	err = s.DB.AutoMigrate(&models.RecordModel{})
	if err != nil {
		return fmt.Errorf("ошибка при миграции - %s", err)
	}
	return nil
}

func (s *Sql) CreateRecord(data *schemas.CreateRecordStruct) error {
	currentRecord := &models.RecordModel{
		Name:       data.Name,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
		Age:        data.Age,
		Gender:     data.Gender,
		Country:    data.Country,
	}

	err := s.DB.Create(currentRecord).Error
	return err
}

func (s *Sql) UpdateRecordByName(data *schemas.UpdateRecordByNameRequestBody) error {
	var currentRecord models.RecordModel
	err := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Name", data.NewName).Error
	return err
}

func (s *Sql) UpdateRecordBySurname(data *schemas.UpdateRecordBySurnameRequestBody) error {
	var currentRecord models.RecordModel
	err := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Surname", data.NewSurname).Error
	return err
}

func (s *Sql) UpdateRecordByAge(data *schemas.UpdateRecordByAgeRequestBody) error {
	var currentRecord models.RecordModel
	err := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Age", data.NewAge).Error
	return err
}

func (s *Sql) DeleteRecord(ID int) error {
	var currentRecord models.RecordModel
	err := s.DB.Where("ID = ?", ID).Delete(&currentRecord).Error
	return err
}

func (s *Sql) GetAllRecords() []models.RecordModel {
	var response []models.RecordModel
	s.DB.Find(&response)
	return response
}
