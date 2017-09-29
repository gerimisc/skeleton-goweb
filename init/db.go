package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Configuration struct for JSON config file
type Configuration struct {
	Hostname         string
	StaticVariable   string
	ConnectionString string
	Username         string
	Password         string
}

func Connect(cfg *Configuration) error {
	db, err := sql.Open("mysql", cfg.Username+":"+cfg.Password+"@/"+cfg.ConnectionString)
	checkErr(err)
	users, err := db.Query("SELECT * FROM users")

	for users.Next() {
		var host string
		var user string
		var password string
		err = users.Scan(&host, &user, &password)
		checkErr(err)

		fmt.Println(host)
	}
	db.Close()

	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
