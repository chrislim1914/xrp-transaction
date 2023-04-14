package models

import (
	"github.com/jinzhu/gorm"
)

type Balances struct {
	*gorm.Model
	UUID      string  `json:"uuid"`
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
	Hold      float64 `json:"hold"`
}
