package models

import (
	"github.com/jinzhu/gorm"
)

type SchemaMigrations struct {
	*gorm.Model
	Version int  `json:"version"`
	Dirty   bool `json:"dirty"`
}
