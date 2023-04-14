package accounts

import (
	"errors"
	"net/http"

	"github.com/chrislim1914/xrp-transaction/api/models"
	"github.com/chrislim1914/xrp-transaction/utility"
	"github.com/google/uuid"
)

func (a *AccountControllerService) GetAccount(request AccountRequest) (models.Accounts, int, error) {
	if err := a.DB.Where("uuid = ?", request.UUID).First(&a.Account).Error; err != nil {
		return *a.Account, http.StatusInternalServerError, err
	}

	return *a.Account, http.StatusOK, nil
}

func (a *AccountControllerService) NewAccount(request AccountRequest) (response CreateAccountResponse, sc int, err error) {
	// validate email request
	checkemail, err := a.verifyEmail(request.Email)
	if !checkemail || err != nil {
		return response, http.StatusInternalServerError, err
	}

	// validate password
	checkpass := utility.VerifyNewPassword(request.Password)
	if !checkpass {
		err = errors.New("malformed password")
		return response, http.StatusInternalServerError, err
	}

	hashpass, err := utility.HashPassword(request.Password)
	if err != nil {
		err = errors.New("hashinf failed")
		return response, http.StatusInternalServerError, err
	}

	a.Account.UUID = uuid.New().String()
	a.Account.AccountName = request.AccountName
	a.Account.Email = request.Email
	a.Account.Password = hashpass
	a.Account.ApiKey = uuid.New().String()
	a.Account.ApiSecret = utility.GenerateSecretKey(*a.Account)

	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return response, http.StatusInternalServerError, err
	}

	if err := tx.Create(&a.Account).Error; err != nil {
		tx.Rollback()
		return response, http.StatusInternalServerError, err
	}

	var balance models.Balances
	balance.UUID = a.Account.UUID

	if err := tx.Create(&balance).Error; err != nil {
		tx.Rollback()
		return response, http.StatusInternalServerError, err
	}

	tx.Commit()

	response.ApiKey = a.Account.ApiKey
	response.ApiSecret = a.Account.ApiSecret
	return response, http.StatusOK, nil
}
