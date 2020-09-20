package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgress diver
)

// Env Variable Names
const (
	psqlUserName = "PSQL_USER_NAME"
	psqlPassword = "PSQL_PASSWORD"
	psqlPort     = "PSQL_PORT"
	psqlDBName   = "PSQL_DB_NAME"

	// Query to create a table if does not exist
	queryCreateTable = `CREATE TABLE IF NOT EXISTS users(
			id serial PRIMARY KEY,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			email VARCHAR(50) UNIQUE NOT NULL,
			status VARCHAR(10) NOT NULL,
			password VARCHAR(50) NOT NULL,
			date_created TIMESTAMP NOT NULL,
			date_updated TIMESTAMP
		);`
)

var (
	// Client db that can be accesed by domain layer
	Client *sql.DB
)

func init() {
	var err error

	// Loading the env variables
	godotenv.Load(".env")

	// Generating the connection string
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv(psqlUserName), os.Getenv(psqlPassword), os.Getenv(psqlDBName), os.Getenv(psqlPort))
	Client, err = sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	_, err = Client.Exec(queryCreateTable)
	if err != nil {
		panic(err)
	}

	log.Println("Database started sucessfully")
}
