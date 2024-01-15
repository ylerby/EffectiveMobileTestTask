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

func (s *Sql) ReadRecordByFullName(name, surname, patronymic string) ([]models.RecordModel, error) {
	var result []models.RecordModel
	err := s.DB.Where("Name = ? AND Surname = ? AND Patronymic = ?", name, surname, patronymic).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Sql) ReadRecordByAge(age int) ([]models.RecordModel, error) {
	var result []models.RecordModel
	err := s.DB.Where("Age = ?", age).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Sql) ReadRecordByGender(gender string) ([]models.RecordModel, error) {
	var result []models.RecordModel
	err := s.DB.Where("Gender = ?", gender).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Sql) ReadRecordByCountry(country string) ([]models.RecordModel, error) {
	var result []models.RecordModel
	err := s.DB.Where("Country = ?", country).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Sql) UpdateRecordByName(data *schemas.UpdateRecordByNameRequestBody) error {
	var currentRecord models.RecordModel
	result := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Name", data.NewName)

	if result.RowsAffected == 0 {
		return fmt.Errorf("значение не найдено")
	}

	return result.Error
}

func (s *Sql) UpdateRecordBySurname(data *schemas.UpdateRecordBySurnameRequestBody) error {
	var currentRecord models.RecordModel
	result := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Surname", data.NewSurname)

	if result.RowsAffected == 0 {
		return fmt.Errorf("значение не найдено")
	}

	return result.Error
}

func (s *Sql) UpdateRecordByPatronymic(data *schemas.UpdateRecordByPatronymicRequestBody) error {
	var currentRecord models.RecordModel
	result := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Patronymic", data.NewPatronymic)

	if result.RowsAffected == 0 {
		return fmt.Errorf("значение не найдено")
	}

	return result.Error
}

func (s *Sql) UpdateRecordByAge(data *schemas.UpdateRecordByAgeRequestBody) error {
	var currentRecord models.RecordModel
	result := s.DB.Where("ID = ?", data.ID).Model(&currentRecord).Update("Age", data.NewAge)

	if result.RowsAffected == 0 {
		return fmt.Errorf("значение не найдено")
	}

	return result.Error
}

func (s *Sql) DeleteRecord(ID int) error {
	var currentRecord models.RecordModel
	result := s.DB.Where("ID = ?", ID).Delete(&currentRecord)

	if result.RowsAffected == 0 {
		return fmt.Errorf("значение не найдено")
	}

	return result.Error
}

func (s *Sql) GetAllRecords() []models.RecordModel {
	var response []models.RecordModel
	s.DB.Find(&response)
	return response
}
