package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

    "gitlab.com/gnufred/logslate/parse/slate"

	_ "github.com/mattn/go-sqlite3"
)

type Dber struct {
	dns  string
	file string
	db   *sql.DB
}

var once sync.Once
var dbInst *Dber

func Singleton() *Dber {
	once.Do(func() {
		dbInst = &Dber{}
	})
	return dbInst
}

func (d *Dber) Create(file string) error {
	d.dns = "file:" + file + "?cache=shared&mode=rw"
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
	sql := `CREATE TABLE logslate (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "address" string,
        "user" string,
        "time" timestamp,
        "method" string,
        "path" string,
        "query" string,
        "proto" string,
        "status" integer,
        "sent" integer,
        "referer" string,
        "agent" string
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

func (d *Dber) InsertLog(e slate.Log) error {
	sql := `INSERT INTO logslate (address,user,time,method,path,query,proto,status,sent,referer,agent) VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	st, err := d.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = st.Exec(
        e.Address,
        e.User,
        e.Time,
        e.Method,
        e.Path,
        e.Query,
        e.Proto,
        e.Status,
        e.Sent,
        e.Referer,
        e.Agent,
    )
	if err != nil {
		return err
	}
	return nil
}

/* ====================
Related documentation
* https://www.sqlite.org/cli.html

Supported types
* https://pkg.go.dev/github.com/mattn/go-sqlite3
+------------------------------+
|go        | sqlite3           |
|----------|-------------------|
|nil       | null              |
|int       | integer           |
|int64     | integer           |
|float64   | float             |
|bool      | integer           |
|[]byte    | blob              |
|string    | text              |
|time.Time | timestamp/datetime|
+------------------------------+

Examples
* https://earthly.dev/blog/golang-sqlite/

*/
