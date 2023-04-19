package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

    "gitlab.com/gnufred/logslate/ngxprs"

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
	sql := `CREATE TABLE logs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "remoteaddr" string,
        "remoteuser" string,
        "timelocal" timestamp,
        "requestmethod" string,
        "requesturl" string,
        "requestprotocol" string,
        "status" integer,
        "bytessent" integer,
        "httpreferer" string,
        "httpuseragent" string,
        "gzipratio" float
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

func (d *Dber) InsertLog(e ngxprs.LogEntry) error {
	sql := `INSERT INTO logs (remoteaddr,remoteuser,timelocal,requestmethod,requesturl,requestprotocol,status,bytessent,httpreferer,httpuseragent,gzipratio) VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	st, err := d.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = st.Exec(
        e.RemoteAddr,
        e.RemoteUser,
        e.TimeLocal,
        e.RequestMethod,
        e.RequestUrl,
        e.RequestProtocol,
        e.Status,
        e.BytesSent,
        e.HttpReferer,
        e.HttpUserAgent,
        e.GzipRatio,
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
