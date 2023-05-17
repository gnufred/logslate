package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"gitlab.com/gnufred/logslate/parse/slate"

	_ "github.com/mattn/go-sqlite3"
)

type Dber struct {
	dns  string
	file string
	db   *sql.DB
	open bool
	tx   *sql.Tx
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
	d.file = file
	d.dns = "file:" + file + "?mode=rw"

	var err error
	if d.file != ":memory:" {
		err = d.createFileIfNotExists()
		if err != nil {
			return fmt.Errorf("unable to create database: %s", err)
		}
	}

	d.db, err = sql.Open("sqlite3", d.dns)
	if err != nil {
		return err
	}
	d.open = true

	d.db.Exec("PRAGMA synchronous = OFF")
	d.db.Exec("PRAGMA journal_mode = OFF")

	err = d.createTables()
	if err != nil {
		return err
	}

	return nil
}

func (d *Dber) createTables() error {
	if !d.open {
		return fmt.Errorf("database closed")
	}
	sql := `CREATE TABLE IF NOT EXISTS logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "ip" string,
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
	d.open = false
	return nil
}

func (d *Dber) createFileIfNotExists() error {
	_, err := os.Stat(d.file)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", d.file)
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

func (d *Dber) Begin() error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	d.tx = tx
	return nil
}

func (d *Dber) Commit() error {
	return d.tx.Commit()
}

func (d *Dber) InsertLog(e slate.Log) error {
	if !d.open {
		return fmt.Errorf("database closed")
	}
	sql := `INSERT INTO logs (ip,user,time,method,path,query,proto,status,sent,referer,agent) VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	st, err := d.tx.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = d.tx.Stmt(st).Exec(
		e.Ip.String(),
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
