package controller

import (
	"Golang/banco"
	"Golang/models"
	"Golang/repository"
	"Golang/responses"
	"Golang/security"
	"Golang/auth" 
	"encoding/json"
	"io"
	"net/http"
	
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, _ := io.ReadAll(r.Body)

	var user models.User
	if err := json.Unmarshal(corpoRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepository(db)
	
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioSalvoNoBanco.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}


	w.Write([]byte(token)) 
	
}