package db_handling

//on imports https://www.golangprograms.com/golang-import-function-from-another-folder.html

import (
	"context"
	"database/sql"
	"fmt"

	"backend_server/entities"

	pg "github.com/go-pg/pg/v10"
	orm "github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
)

type Account = entities.Account

func Set_up_db() {
	db := ConnectDBORM()
	err := set_up_tables(db)
	if err != nil {
		panic(err)
	}
}

func set_up_tables(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Account)(nil),
		(*Transaction)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ConnectDBORM() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "db_connect",
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return db

}

func ConnectDbSQL() (db *sql.DB) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "db_connect"

	//zeigt es ist sowohl über SQL Injection als auch ORM möglich
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return
}
