package repo

import "gorm.io/gorm"

type Repo struct {
	orm *gorm.DB
}

func (r *Repo ) Transcation(execute func(tx *gorm.DB) error) error  {
	return r.orm.Transaction(execute)
}

func (r *Repo ) Begin() *gorm.DB {
	return r.orm.Begin()
}