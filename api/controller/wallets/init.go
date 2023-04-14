package wallets

import (
	"errors"

	"github.com/chrislim1914/xrp-transaction/api/models"
	"github.com/chrislim1914/xrp-transaction/bootstrap"
	"github.com/chrislim1914/xrp-transaction/database"
	"gorm.io/gorm"
)

type WalletController interface {
	GetClientWallet(uuid string) (models.Wallets, error)
	NewInternalWallet(r NewWalletRequest) (response NewWalletResponse, sc int, err error)
}

type WalletControllerService struct {
	DB     *gorm.DB
	Wallet *models.Wallets
}

type NewWalletRequest struct {
	Address        string `json:"address"`
	DestinationTag string `json:"destination_tag,omitempty"`
}

type NewWalletResponse struct {
	Message string `json:"message"`
}

func NewWalletController() WalletController {
	config, _ := bootstrap.LoadConfig(".")
	db, _ := database.ConnectDB(&config)
	return &WalletControllerService{
		DB:     db,
		Wallet: &models.Wallets{},
	}
}

func (w *WalletControllerService) verifyWallet(address, tag string) (bool, error) {
	if err := w.DB.Where("address = ? AND destination_tag = ?", address, tag).First(&w.Wallet).Error; err != nil {
		return true, nil
	}
	err := errors.New("address already exist...")
	return false, err
}
