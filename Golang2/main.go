package main

import (
	"Golang/banco"
	"Golang/router"
	"log"
	"net/http"
)

func main() {
	db, err := banco.Conectar()
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados: ", err)
	}
	db.Close() 

	log.Println("Conexão com banco de dados verificada!")

	r := router.New()
	const addr = ":8080"
		log.Printf("Starting server on %s\n", addr) 
	
	log.Fatal(http.ListenAndServe(addr, r))
}