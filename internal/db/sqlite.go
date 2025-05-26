package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dataDir          = "data"
	sqliteDbFilename = "wizard.db"
)

func InitSqliteDB() (*sql.DB, error) {
	if err := ensureDataDir(); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dataDir, sqliteDbFilename)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := executeSQLSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

func ensureDataDir() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return os.MkdirAll(dataDir, 0755)
	}
	return nil
}

func readSQLFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func executeSQLSchema(db *sql.DB) error {
	schema, err := readSQLFile("schemas/001_create_scripts.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(schema)
	return err
}

func CloseDB(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
