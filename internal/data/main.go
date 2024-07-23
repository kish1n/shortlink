package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "gitlab.com/distributed_lab/kit/pgdb"
	"log"
	"os"
)

type MasterQ interface {
	New() MasterQ

	Nonce() LinksQ

	Transaction(fn func(db MasterQ) error) error
}

func InitDB() (*sql.DB, error) {

	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbName == "" || dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("missing one or more environment variables: POSTGRES_DB=%s, POSTGRES_USER=%s, POSTGRES_PASSWORD=%s, DB_HOST=%s, DB_PORT=%s",
			dbName, dbUser, dbPassword, dbHost, dbPort)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	fmt.Println("connStr: ", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	log.Println("Connected to the database successfully")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS links (
		original TEXT NOT NULL,
		shortened TEXT PRIMARY KEY
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	log.Println("Table 'links' checked/created successfully")

	return db, nil
}
