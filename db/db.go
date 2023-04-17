package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Dber struct {
	dns  string
	file string
	conn *sql.DB
}

var dber = Dber{}

func (d *Dber) Create(dns string, file string) error {
	d.dns = dns
	d.file = file

	err := createFileIfNotExists(d.file)
	if err != nil {
		return err
	}

	d.conn, err = sql.Open("sqlite3", d.file)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dber) Close() error {
	err := d.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func createFileIfNotExists(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) == false {
		return nil
	}

	fileHandle, err := os.Create(file)
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
