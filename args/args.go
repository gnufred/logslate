package args

import "flag"

type Args struct {
	Db string
}

func Parse() Args {
	dbFlag := flag.String("db", "", "Save path for the SQLite database")
	flag.Parse()

	return Args{
		Db: string(*dbFlag),
	}
}
