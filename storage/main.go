package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE test (
			id int,
			balance_decimal decimal(16, 8),
			balance_float float(16, 8),
			PRIMARY KEY (id)
		);
	`)

	_, err = db.Exec("INSERT INTO test (id, balance_decimal, balance_float) VALUES (1, 1.1, 1.1)")
	_, err = db.Exec(`
		UPDATE test 
			SET 
			    balance_decimal = balance_decimal + 1.2 , 
			    balance_float   = balance_float + 1.2 
			WHERE id = 1;
	`)
}
