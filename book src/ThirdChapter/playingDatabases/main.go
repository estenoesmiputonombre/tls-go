package main

import (
	"database/sql"
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
	//Driver ... Used to manage connections and data
	Driver *sql.DB
)

func main() {

}
