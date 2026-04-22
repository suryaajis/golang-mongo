package models

type Book struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
}
