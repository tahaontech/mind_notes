package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Database *sql.DB
}

type PhotoDetail struct {
	Name       string
	TimeTaken  int
	LastViewed int
}

func InitDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "data/app.db")
	if err != nil {
		return nil, err
	}

	d := &DB{
		Database: db,
	}

	initStmt := `
	CREATE TABLE IF NOT EXISTS node (id string unique not null PRIMARY KEY, label text, root bool, positionX float, positionY float, rootId string );
	CREATE TABLE IF NOT EXISTS edge (id string unique not null PRIMARY KEY, sourceId string,  targetId string , rootId string);
	CREATE TABLE IF NOT EXISTS document (id string unique not null PRIMARY KEY, nodeId string, data text );
	`

	_, err = d.Database.Exec(initStmt)
	if err != nil {
		return nil, err
	}

	return d, nil
}
