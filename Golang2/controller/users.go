package controller

import (
	"Golang/banco"
	"Golang/models"
	"Golang/repository"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	corpoRequest, _ := io.ReadAll(r.Body)

	var user models.User
	if err := json.Unmarshal(corpoRequest, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := user.Prepare("cadastro"); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()


	repositorio := repository.NewUsersRepository(db)
	
	insertedID, err := repositorio.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário criado com ID: " + strconv.FormatUint(insertedID, 10)))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, _ := strconv.ParseUint(parametros["userID"], 10, 64)

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()


	repositorio := repository.NewUsersRepository(db)

	user, err := repositorio.SearchByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, _ := strconv.ParseUint(parametros["userID"], 10, 64)

	corpoRequest, _ := io.ReadAll(r.Body)
	var user models.User
	if err := json.Unmarshal(corpoRequest, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := user.Prepare("edicao"); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepository(db)

	if err := repositorio.Update(userID, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, _ := strconv.ParseUint(parametros["userID"], 10, 64)

	
	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	
	repositorio := repository.NewUsersRepository(db)

	if err := repositorio.Delete(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	db, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	
	repositorio := repository.NewUsersRepository(db)

	users, err := repositorio.SearchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}