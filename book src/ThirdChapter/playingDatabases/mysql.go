package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

var (
	//DBHost ... Address of the database to connect to it
	DBHost = os.Getenv("MYSQL_ROOT_ADDR")
	//DBPort ... Port to connect to the database
	DBPort = os.Getenv("MYSQL_ROOT_PORT")
	//DBUser ... User to connect to the database
	DBUser = os.Getenv("MYSQL_ROOT_USER")
	//DBPassword ... Password to connect to the database
	DBPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	//DBDbase ... Database name to connect
	DBDbase = os.Getenv("MYSQL_ROOT_DATABASE_NAME")
	//Driver ... Driver of mysql
	//Driver *sql.DB
)

var (
	//ErrNegativeValue ... Error parsing index less or equal than 0
	ErrNegativeValue = errors.New("Error parsing index less than 0")
)

//ConnectDB ... Returns a fresh instance of a driver to connect to mysql
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", GetStringToConnect())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}

//GetStringToConnect ... Return string to connect to sql
func GetStringToConnect() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPassword, DBHost, DBPort, DBDbase) //username:password@tcp(address:port)/database
}

//GetNEmployers ... Return n employers struct
func GetNEmployers(n int) ([]Employer, error) {
	if n <= 0 {
		return nil, ErrNegativeValue
	}
	driver, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	err = driver.Ping()
	if err != nil {
		return nil, err
	}
	rows, err := driver.QueryRow("SELECT e.name, e.surname, l.name FROM employers e INNER JOIN languages l ON e.id_language = l.id LIMIT ?;", n).Scan()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
