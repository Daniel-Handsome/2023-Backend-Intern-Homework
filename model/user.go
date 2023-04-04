package model

type User struct {
	BaseModel
	Uuid           string
	Token          string
	Name           string
	Email          string
	PageLinkedList []PageLinkedList `gorm:"foreignKey:UserId;references:id"`
}
