package models

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name string `gorm:"column: name; type: varchar(5);" json:"Name"`
}

func (table *TestModel) TableName() string {
	return "test"
}
