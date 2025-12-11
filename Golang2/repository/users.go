package repository

import (
	"Golang/models"
	"database/sql"
	"fmt" 
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (nome, email, password, cpf) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(user.Nome, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (u users) SearchByID(ID uint64) (models.User, error) {
	linha, err := u.db.Query("select id, nome, email, cpf, criadoem from users where id = ?", ID)
	if err != nil {
		return models.User{}, err
	}
	defer linha.Close()

	var user models.User

	if linha.Next() {
		if err = linha.Scan(&user.ID, &user.Nome, &user.Email, &user.CPF, &user.CriadoEm); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, fmt.Errorf("usuário não encontrado")
	}

	return user, nil
}

func (u users) SearchAll() ([]models.User, error) {
	linhas, err := u.db.Query("select id, nome, email, cpf, criadoem from users")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.User

	for linhas.Next() {
		var user models.User
		if err = linhas.Scan(&user.ID, &user.Nome, &user.Email, &user.CPF, &user.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, user)
	}

	return usuarios, nil
}

func (u users) Update(ID uint64, user models.User) error {
	statement, err := u.db.Prepare("update users set nome = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Nome, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (u users) Delete(ID uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}


func (u users) BuscarPorEmail(email string) (models.User, error) {
	
	linha, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer linha.Close()

	var user models.User

	if linha.Next() {
		if err = linha.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, fmt.Errorf("usuário não encontrado")
	}

	return user, nil
}
