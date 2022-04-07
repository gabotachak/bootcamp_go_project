package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

/**
	function load database
**/
func Init(environment string) (*sql.DB, error) {
	switch environment {
    case "test":
      err := godotenv.Load("../../../.env")
      if err != nil {
        log.Fatal("Error loading .env file")
      }

    case "unit_test":
      err := godotenv.Load("../../.env")
      if err != nil {
        log.Fatal("Error loading .env file")
      }
    default:
      err := godotenv.Load(".env")
      if err != nil {
        log.Fatal("Error loading .env file")
      }
	}

	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if environment == "test" || environment == "unit_test" {
		dbName = os.Getenv("DB_NAME_TEST")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
	StorageDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	// Check that the database is available and accessible
	e := StorageDB.Ping()
	if e != nil {
		fmt.Println("error en el ping", e)
	}
	return StorageDB, nil
}
