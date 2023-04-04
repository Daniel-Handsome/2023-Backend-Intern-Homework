package model

import "github.com/lib/pq"

type PageNode struct {
	BaseModel
	Uuid       string
	ArticleIds pq.Int64Array `gorm:"type:bigint[]"`
	Previous   string
	Next       string
	//Articles   []Article `gorm:"foreignKey:ID;references:ArticleIds"`
}
