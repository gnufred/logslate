package db

type Dber struct {
	dns  string
	file string
}

var dber = Dber{}

func (d *Dber) New(dns string, file string) {
	d.dns = dns
	d.file = file
}
