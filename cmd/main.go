package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IpAddress string `json:"ip_address"`
}

func main() {
	ConnectDatabase()

	e := echo.New()

	group := e.Group("/api")

}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./db/accounts.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
