/* Figure out how to use SQLite3 connection directly, and not as a database/sql driver
   https://github.com/mattn/go-sqlite3
   https://pkg.go.dev/github.com/mattn/go-sqlite3
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// init
	timer := time.Now()

	dbFile := "test.db"
	err := CreateFile(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer DestroyFile(dbFile)

	db, err := OpenDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDB(db)

	err = CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v => init\n", time.Since(timer))

	testStr := "this is a text line with text in it that is also a string"

	// 1000 single inserts
	timer = time.Now()
	InsertNumRows(db, 1000, testStr)
	fmt.Printf("%v => 1000 single inserts\n", time.Since(timer))

	// 1000 transact inserts
	timer = time.Now()
	TransactNumRows(db, 1000, testStr)
	fmt.Printf("%v => 1000 transact inserts\n", time.Since(timer))

	// 1000 single inserts
	db.Exec("PRAGMA synchronous = OFF")
	timer = time.Now()
	InsertNumRows(db, 1000, testStr)
	fmt.Printf("%v => 1000 single inserts (sync off)\n", time.Since(timer))

	// 1000 single inserts
	db.Exec("PRAGMA journal_mode = MEMORY")
	timer = time.Now()
	InsertNumRows(db, 1000, testStr)
	fmt.Printf("%v => 1000 single inserts (sync off, journal mem)\n", time.Since(timer))

	// 1000 transact inserts
	timer = time.Now()
	TransactNumRows(db, 1000, testStr)
	fmt.Printf("%v => 1000 transact inserts (sync off, journal mem)\n", time.Since(timer))
}

func TransactNumRows(db *sql.DB, num int, row string) {
	tx, _ := db.Begin()
	for i := 0; i < num; i++ {
		sql := `INSERT INTO test (row) VALUES (?);`
		_, err := tx.Query(sql, row)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func InsertNumRows(db *sql.DB, num int, row string) {
	for i := 0; i < num; i++ {
		err := InsertSingleRow(db, row)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertSingleRow(db *sql.DB, row string) error {
	sql := `INSERT INTO test (row) VALUES (?);`
	st, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = st.Exec(row)
	if err != nil {
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) error {
	sql := `CREATE TABLE test ("row" text);`
	st, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = st.Exec()
	if err != nil {
		return err
	}
	return nil
}

func OpenDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return db, err
	}
	return db, err
}

func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) != true {
		return fmt.Errorf("File %s already exists", file)
	}

	fileHandle, err := os.Create(file)
	if err != nil {
		return err
	}
	fileHandle.Close()

	return nil
}

func DestroyFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}
