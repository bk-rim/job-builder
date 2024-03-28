package database

import (
	_ "github.com/lib/pq"

	"database/sql"
	"os"
)

var DB *sql.DB

func InitDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslmode := os.Getenv("DB_SSL_MODE")

	var dbHost string
	if os.Getenv("DOCKER_ENV") == "true" {
		dbHost = os.Getenv("DB_HOST_DOCKER")
	} else {
		dbHost = os.Getenv("DB_HOST")
	}

	connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSslmode

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	DB = db
	migrate()

}

func migrate() {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS jobs (
			id SERIAL PRIMARY KEY,
			name VARCHAR(30) NOT NULL UNIQUE,
			type VARCHAR(30) NOT NULL,
			frequency VARCHAR(30) NOT NULL,
			created_on DATE NOT NULL,
			status VARCHAR(30) NOT NULL,
			executed_on DATE NULL,
			webhook_slack VARCHAR(100) NOT NULL
		);
	`
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
