package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chrislim1914/xrp-transaction/api/controller/wallets"
)

func NewInternalWallet(w http.ResponseWriter, r *http.Request) {
	walletsrv := wallets.NewWalletController()
	var request wallets.NewWalletRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, sc, err := walletsrv.NewInternalWallet(request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), sc)
		return
	}
	jsonResponseHandler(w, response, sc)
}
