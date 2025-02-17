package config

import (
	"database/sql"
	"fmt"
	"time"
)

func Wait() {

	dsn := "root:pass@tcp(mysql-container:3306)/db"
	for {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Waiting for MySQL to be ready...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			fmt.Println("MySQL is up and running!")
			db.Close()
			break
		}

		db.Close()
		fmt.Println("Waiting for MySQL to be ready...")
		time.Sleep(5 * time.Second)
	}
}
