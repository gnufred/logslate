package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"gitlab.com/gnufred/logslate/log"
	"gitlab.com/gnufred/logslate/test"
	nginx "gitlab.com/gnufred/nxparse"
)

func TestCreate(t *testing.T) {
	createDB(t)
	recreateDB(t)
}

func createDB(t *testing.T) {
	file := "testdata/test.db"
	os.Remove(file)
	dbi := Singleton()
	err := dbi.Create(file)
	if err != nil {
		t.Fatalf("Create() failed: %s", err)
	}
	dbi.ApplyDefaults()
}

func recreateDB(t *testing.T) {
	file := "testdata/test.db"
	dbi := Singleton()
	err := dbi.Create(file)
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("unexpecte error, got %s", err)
	}
}

func TestInsert(t *testing.T) {
	file, err := os.Open("testdata/ahrefs.log")
	if err != nil {
		t.Fatalf("Opening test data file failed: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dbi := Singleton()
	dbi.Begin()
	defer dbi.Commit()

	for scanner.Scan() {
		nx, _ := nginx.Parse(nginx.Combined, scanner.Text())
		if err != nil {
			t.Errorf("nginx.Parsed() failed: %s", err)
		}
		sl, err := log.NginxToSlate(nx)
		if err != nil {
			t.Error(err)
		}
		err = dbi.Insert(sl)
		if err != nil {
			t.Errorf("Insert() failed: %s", err)
		}
	}
}

func TestClose(t *testing.T) {
	closeDB(t)
	recloseDB(t)
	closeCreateError(t)
}

func closeDB(t *testing.T) {
	dbi := Singleton()
	err := dbi.Close()
	if err != nil {
		t.Errorf("Close() failed: %s", err)
	}
}

func recloseDB(t *testing.T) {
	dbi := Singleton()
	err := dbi.Close()
	if err != nil {
		t.Errorf("Close() failed: %s", err)
	}
}

func closeCreateError(t *testing.T) {
	dbi := Singleton()
	err := dbi.createTables()
	test.ErrorExpected(t, err, "closed")
}

func TestDestroy(t *testing.T) {
	destroyDB(t)
	redestroyDB(t)
}

func destroyDB(t *testing.T) {
	dbi := Singleton()
	err := dbi.Destroy()
	if err != nil {
		t.Errorf("Destroy() failed: %s", err)
	}
}

func redestroyDB(t *testing.T) {
	dbi := Singleton()
	err := dbi.Destroy()
	if !strings.Contains(err.Error(), "no such file") {
		t.Errorf("unexpedted error, got %s", err)
	}
}

// === benchmarks ===

func BenchmarkDbInsertBulk(b *testing.B) {
	benchDbInsert(b, 200, []int{1, 3, 5, 7, 9, 15, 20, 30}, []string{})
}

func BenchmarkDbWithDefaults(b *testing.B) {
	benchDbInsert(b, 200, []int{1, 3, 5, 7, 9, 15, 20, 30}, []string{
		"synchronous = off",
		"journal_mode = off",
		"locking_mode = exclusive",
		"transaction_mode = exclusive",
	})
}

func benchDbInsert(b *testing.B, rows int, prepare []int, pragmas []string) {
	// The range is used to check bulk insert performance differences for every "v" insert.
	// Seems to me like we get the best results when we prepare "7" inserts.
	for _, v := range prepare {
		b.Run(fmt.Sprintf("DbInsert-%d", v), func(b *testing.B) {
			dbi := Singleton()
			lineCombined := "216.245.221.84 - - [27/Apr/2023:09:49:32 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://example.com/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""
			nx, _ := nginx.Parse(nginx.Combined, lineCombined)
			sl, _ := log.NginxToSlate(nx)

			// benchmark loop
			for i := 0; i < b.N; i++ {
				err := dbi.Create(":memory:")
				if err != nil {
					b.Fatal(err)
				}
				dbi.Begin()
				for _, p := range pragmas {
					dbi.Pragma(p)
				}

				// bulk insert loop
				for j := range rows {
					err = dbi.Insert(sl)

					// Commit every "v" lines.
					if j%v == 0 {
						dbi.Commit()
						dbi.Begin()
					}

					if err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}
