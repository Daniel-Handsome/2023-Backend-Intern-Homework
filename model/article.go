package model

type Article struct {
	BaseModel
	Uuid    string
	Title   string
	Content string
	Sort    int8
}
