package models

import (
	"github.com/jinzhu/gorm"
)

type Wallets struct {
	*gorm.Model
	UUID           string `json:"uuid"`
	Address        string `json:"address"`
	DestinationTag string `json:"destination_tag"`
}
