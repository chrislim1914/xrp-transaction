package middleware

import (
	"net/http"
	"strings"

	"github.com/chrislim1914/xrp-transaction/api/controller/accounts"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get api key and secret
		apikey := r.Header.Get("X-API-KEY")
		apisecret := r.Header.Get("X-API-SECRET")

		// get account data
		acc := accounts.NewAccountController()
		data := acc.GetAPIData(apikey)
		if !strings.EqualFold(apisecret, data) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
