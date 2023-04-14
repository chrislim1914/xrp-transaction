package wallets

import (
	"net/http"

	"github.com/chrislim1914/xrp-transaction/api/models"
)

func (w *WalletControllerService) GetClientWallet(uuid string) (models.Wallets, error) {
	if err := w.DB.Where("uuid = ?", uuid).First(&w.Wallet).Error; err != nil {
		return *w.Wallet, err
	}

	return *w.Wallet, nil
}

func (w *WalletControllerService) NewInternalWallet(r NewWalletRequest) (response NewWalletResponse, sc int, err error) {
	// verify wallet
	isgood, err := w.verifyWallet(r.Address, r.DestinationTag)
	if !isgood || err != nil {
		return response, http.StatusInternalServerError, err
	}

	w.Wallet.Address = r.Address
	w.Wallet.DestinationTag = r.DestinationTag
	tx := w.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return response, http.StatusInternalServerError, err
	}

	if err := tx.Create(&w.Wallet).Error; err != nil {
		tx.Rollback()
		return response, http.StatusInternalServerError, err
	}

	tx.Commit()
	response.Message = "New Internal Wallet save!"
	return response, http.StatusOK, nil
}
