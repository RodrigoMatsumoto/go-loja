package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "root:root@/alura_loja"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
