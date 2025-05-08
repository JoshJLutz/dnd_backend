package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "postgresql-test-1.cbmoywcegkxx.us-east-2.rds.amazonaws.com"
	port = 5432
	user = "jjl_aws"
	password = "joggles199"
	dbname = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("%s:%s@%s:%d/%s?sslmode=require", user, password, host, port, dbname)
	fmt.Println(psqlconn)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected :)")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}