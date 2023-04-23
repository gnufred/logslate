package args

import "flag"

type Args struct {
	Db string
    V bool
    Vv bool
    Vvv bool
}

var arguments = Args{}

func Parse() {
	dbFlag := flag.String("db", "", "Save path for the SQLite database")
    vFlag := flag.Bool("v", false, "Debug: verbose")
    vvFlag := flag.Bool("vv", false, "Debug: very verbose")
    vvvFlag := flag.Bool("vvv", false, "Debug: very very verbose")
	flag.Parse()

	arguments = Args{
		Db: string(*dbFlag),
        V: bool(*vFlag),
        Vv: bool(*vvFlag),
        Vvv: bool(*vvvFlag),
    }
}

func Get() Args {
    return arguments
}
