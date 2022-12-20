package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/taglyscostacurta/api-postgresql/configs"
)

func OpenConnection() (*sql.DB, error) {
	Conf := configs.GetDb()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Conf.Host, Conf.Port, Conf.User, Conf.Pass, Conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
