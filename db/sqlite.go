package db

import (
	"database/sql"
	"log"
	"time"

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

func (d *DB) InitDB() {
	db, err := sql.Open("sqlite3", "./data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	d.Database = db

	initStmt := `
	CREATE TABLE IF NOT EXISTS PHOTOS (FILENAME string unique not null primary key, LAST_VIEWED int, TIME_TAKEN int )
	`

	_, err = d.Database.Exec(initStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, initStmt)
	}
}

func (d *DB) AddPhoto(name string, timeTaken int) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO PHOTOS(FILENAME, LAST_VIEWED, TIME_TAKEN) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, 0, timeTaken)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) UpdatePhoto(name string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE PHOTOS SET LAST_VIEWED=? WHERE FILENAME=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Unix(), name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) GetNextSlidseshowPhoto() (PhotoDetail, error) {
	rows, err := d.Database.Query("select FILENAME, LAST_VIEWED, TIME_TAKEN FROM PHOTOS ORDER BY LAST_VIEWED LIMIT 1")
	if err != nil {
		return PhotoDetail{}, err
	}
	defer rows.Close()

	for rows.Next() {
		details := PhotoDetail{}
		err := rows.Scan(&details.Name, &details.LastViewed, &details.TimeTaken)
		if err != nil {
			return PhotoDetail{}, err
		} else {
			return details, nil
		}
	}

	err = rows.Err()
	return PhotoDetail{}, err
}
