package model

import "fmt"

const PAGE_LIMIT int = 2

type Article struct {
	BaseModel
	Uuid    string
	Title   string
	Content string
	Sort    int8
}

type OrderColumn int8

// Enum value maps for orderColumn.
var (
	OrderColumn_name = map[int8]string{
		0: "id",
		1: "created_at",
		2: "updated_at",
	}
	OrderColumn_value = map[string]int8{
		"id":         0,
		"created_at": 1,
		"updated_at": 2,
	}
)

const (
	Id OrderColumn = iota
	CreateAt
	UpdateAt
)

func (o OrderColumn) Enum() *OrderColumn {
	p := new(OrderColumn)
	*p = o
	return p
}

func (o OrderColumn) Int8() int8 {
	return int8(o)
}

func (o OrderColumn) GetName() (string, error) {
	value, ok := OrderColumn_name[o.Int8()]
	if !ok {
		return "", fmt.Errorf("not found OrderColumn_name")
	}

	return value, nil
}
