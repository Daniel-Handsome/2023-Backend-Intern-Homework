package filter

import "gorm.io/gorm"

type NameFilter struct{}

func (f *NameFilter) ApplyFilter(query *gorm.DB, value interface{}) *gorm.DB {
	name, ok := value.(string)
	if !ok {
		return query
	}
	return query.Where("name = ?", name)
}
