package args

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	args := Parse()
	if args.Db != "" {
		log.Fatal("db flag is not an empty string")
	}
}
