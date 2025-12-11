package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	// Banco:   golang_db
	stringConexao := "root:xcaralhA12!@tcp(127.0.0.1:3306)/bancogo?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
