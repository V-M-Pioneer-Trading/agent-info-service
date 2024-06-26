package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetUpDatabase() *sql.DB {

	log.Default().Printf("Establishing connection to MySql DB...")

	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:example@(mysql:3306)/vnm-agent-db?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	log.Default().Printf("Connection to DB is established.")

	//if err := db.Ping(); err != nil {
	//	log.Fatal(err)
	//}
	//log.Default().Printf("Database is ready.")

	return db
}
