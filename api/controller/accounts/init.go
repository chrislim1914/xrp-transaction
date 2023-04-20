package accounts

import (
	"errors"
	"net/mail"

	"github.com/chrislim1914/xrp-transaction/api/models"
	"github.com/chrislim1914/xrp-transaction/bootstrap"
	"github.com/chrislim1914/xrp-transaction/database"
	"gorm.io/gorm"
)

type AccountController interface {
	GetAccount(request AccountRequest) (models.Accounts, int, error)
	NewAccount(request AccountRequest) (response CreateAccountResponse, sc int, err error)
	GetAPIData(key string) string
}

type AccountRequest struct {
	UUID        string `json:"uuid"`
	AccountName string `json:"account_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type CreateAccountResponse struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type AccountControllerService struct {
	DB      *gorm.DB
	Account *models.Accounts
}

func NewAccountController() AccountController {
	config, _ := bootstrap.LoadConfig(".")
	db, _ := database.ConnectDB(&config)
	return &AccountControllerService{
		DB:      db,
		Account: &models.Accounts{},
	}
}

func (a *AccountControllerService) verifyEmail(email string) (bool, error) {
	// check email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}

	// check if email exist
	if err := a.DB.Where("email = ?", email).First(&a.Account).Error; err != nil {
		return true, nil
	}
	err = errors.New("email already exist...")
	return false, err
}
