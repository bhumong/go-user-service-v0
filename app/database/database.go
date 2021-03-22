package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/bhumong/go-user-service-v0/app/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("error parsing str to int")
	}
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME")))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Connection Opened to Database")
	Up()
	return nil
}
