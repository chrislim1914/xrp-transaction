package route

import (
	"github.com/chrislim1914/xrp-transaction/api/handler"
	cmw "github.com/chrislim1914/xrp-transaction/api/middleware"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	// gorilla/mux router
	r := mux.NewRouter()

	// use the logger middleware on complete router
	r.Use(cmw.LogMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	s := r.PathPrefix("/api/v1").Subrouter()

	// serving routes
	// test routes
	s.HandleFunc("/", handler.TestHandler).Methods("GET")

	// accounts
	s.HandleFunc("/accounts/{uuid}", handler.GetAccount).Methods("GET")
	s.HandleFunc("/accounts/create", handler.NewAccount).Methods("POST")

	// wallets
	// TODO: need auth middleware to get requestee indentity
	s.HandleFunc("/wallets/create", handler.NewInternalWallet).Methods("POST")

	return r
}
