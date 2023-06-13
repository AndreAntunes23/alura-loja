package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConecDB() *sql.DB {
	conexao := "host=localhost user=postgres database=loja_alura port=5432 password=Andrelfa_23 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
