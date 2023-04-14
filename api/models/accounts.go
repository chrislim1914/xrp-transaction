package models

import (
	"github.com/jinzhu/gorm"
)

type Accounts struct {
	*gorm.Model
	UUID        string `json:"uuid"`
	AccountName string `json:"account_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ApiKey      string `json:"api_key"`
	ApiSecret   string `json:"api_secret"`
}
