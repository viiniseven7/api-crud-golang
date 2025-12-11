package models

import (
	"Golang/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CPF      string    `json:"cpf,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}


	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Nome == "" {
		return errors.New("o campo nome é obrigatório")
	}
	if u.Email == "" {
		return errors.New("o campo email é obrigatório")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o email inserido é inválido")
	}
	if step == "cadastro" && u.Password == "" {
		return errors.New("o campo senha é obrigatório")
	}
	if u.CPF == "" || !isCPFValid(u.CPF) {
		return errors.New("o CPF inserido é inválido")
	}
	return nil 
}

func (u *User) format(step string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)

	if step == "cadastro" {
		senhaComHash, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(senhaComHash)
	}
	return nil
}

func isCPFValid(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	return len(cpf) == 11
}

