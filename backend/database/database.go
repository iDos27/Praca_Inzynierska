package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "restaurant_user"
	dbPassword := "restaurant_pass"
	dbName := "restaurant_db"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Błąd przy otwieraniu połączenia z bazą danych: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Nie można połączyć się z bazą danych: ", err)
	}

	fmt.Println("Połączenie z bazą zostało nawiązane")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Połączenie zostało zamknięte")
	}
}
