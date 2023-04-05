package model

type PageLinkedListType int8

const (
	RECOMMEND PageLinkedListType = iota
)

type PageLinkedList struct {
	BaseModel
	Uuid   string
	UserId int64
	Type   int8
	Head   string
}
