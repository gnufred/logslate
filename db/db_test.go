package db

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	nginx "gitlab.com/gnufred/logslate/nxparse"
	"gitlab.com/gnufred/logslate/parse/slate"
)

func TestCreate(t *testing.T) {
	file := "testdata/test.db"
	db := Singleton()
	err := db.Create(file)
	if err != nil {
		t.Fatalf("Create() failed: %s", err)
	}
}

func TestInsertLog(t *testing.T) {
	file, err := os.Open("testdata/ahrefs.log")
	if err != nil {
		t.Fatalf("Opening test data file failed: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	db := Singleton()
	db.Begin()
	defer db.Commit()

	for scanner.Scan() {
		nx, _ := nginx.Parse(nginx.Combined, scanner.Text())
		if err != nil {
			t.Fatalf("nginx.Parsed() failed: %s", err)
		}
		sl := slate.ParseNginx(nx)
		err = db.InsertLog(sl)
		if err != nil {
			t.Fatalf("InsertLog() failed: %s", err)
		}
	}
}

func TestClose(t *testing.T) {
	db := Singleton()
	err := db.Close()
	if err != nil {
		t.Fatalf("Close() failed: %s", err)
	}
}

func TestDestroy(t *testing.T) {
	db := Singleton()
	err := db.Destroy()
	if err != nil {
		t.Fatalf("Destroy() failed: %s", err)
	}
}

// === benchmarks ===

func BenchmarkDbInsert(b *testing.B) {
	// The range is used to check bulk insert performance differences for every "v" insert.
	// Seems to me like we get the best results when we prepare "7" inserts.
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 30, 50} {
		b.Run(fmt.Sprintf("DbInsert-%d", v), func(b *testing.B) {
			db := Singleton()
			lineCombined := "216.245.221.84 - - [27/Apr/2023:09:49:32 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://example.com/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""
			nx, _ := nginx.Parse(nginx.Combined, lineCombined)
			sl := slate.ParseNginx(nx)

			for i := 0; i < b.N; i++ {
				err := db.Create(":memory:")
				if err != nil {
					b.Fatal(err)
				}
				db.Begin()
				for j := 0; j < 200; j++ {
					err = db.InsertLog(sl)

					// Commit every "v" lines.
					if j%v == 0 {
						db.Commit()
						db.Begin()
					}

					if err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}
