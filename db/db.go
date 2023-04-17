package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Dber struct {
	dns  string
	file string
	db   *sql.DB
}

var dber = Dber{}

func (d *Dber) Create(dns string, file string) error {
	d.dns = dns
	d.file = file

	err := d.createFileIfNotExists()
	if err != nil {
		log.Fatalf("Fatal error: Unable to create database file. Error: %s.", err)
	}

	d.db, err = sql.Open("sqlite3", d.file)
	if err != nil {
		return err
	}

	d.createTables()

	return nil
}

func (d *Dber) createTables() error {
	sql := `CREATE TABLE logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"raw" TEXT
	);`

	st, err := d.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = st.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dber) Close() error {
	err := d.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d *Dber) createFileIfNotExists() error {
	_, err := os.Stat(d.file)
	if os.IsNotExist(err) == false {
		return fmt.Errorf("File %s already exists", d.file)
	}

	fileHandle, err := os.Create(d.file)
	if err != nil {
		return err
	}
	fileHandle.Close()

	return nil
}

func (d *Dber) Destroy() error {
	err := os.Remove(d.file)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dber) InsertLog(raw string) error {
	sql := `INSERT INTO logs (raw) VALUES (?)`
	st, err := d.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = st.Exec(raw)
	if err != nil {
		return err
	}
	return nil
}
