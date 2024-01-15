package models

type RecordModel struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(255)"`
	Surname    string `gorm:"type:varchar(255)"`
	Patronymic string `gorm:"type:varchar(255)"`
	Age        int    `gorm:"type:INTEGER"`
	Gender     string `gorm:"type:varchar(255)"`
	Country    string `gorm:"type:varchar(255)"`
}
