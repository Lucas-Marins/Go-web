package db

import (
	_"github.com/lib/pq"
	"database/sql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=1234567 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}else{
		return db
	}
}