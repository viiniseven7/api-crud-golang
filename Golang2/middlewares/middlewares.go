package middlewares

import (
	"Golang/auth"
	"Golang/responses"
	"net/http"
)


func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if err := auth.ValidarToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}

		
		proximaFuncao(w, r)
	}
}