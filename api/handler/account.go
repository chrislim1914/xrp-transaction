package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/chrislim1914/xrp-transaction/api/controller/accounts"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	account := accounts.NewAccountController()
	var request accounts.AccountRequest
	path := strings.Split(r.URL.Path, "/")
	request.UUID = (path[len(path)-1])
	response, sc, err := account.GetAccount(request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), sc)
		return
	}
	jsonResponseHandler(w, response, sc)
}

func NewAccount(w http.ResponseWriter, r *http.Request) {
	account := accounts.NewAccountController()
	var request accounts.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, sc, err := account.NewAccount(request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), sc)
		return
	}
	jsonResponseHandler(w, response, sc)
}
