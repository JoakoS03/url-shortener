package models

type URL struct {
	ID          int    `gorm:"primarykey"`
	OriginalURL string `gorm:"type:text;not null"`
	ShortURL    string `gorm:"type:varchar(10);unique;not null"`
}
