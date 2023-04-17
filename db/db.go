package db

import (
	"log"
	"os"
)

type Dber struct {
	dns  string
	file string
}

var dber = Dber{}

func (d *Dber) Create(dns string, file string) {
	d.dns = dns
	d.file = file

	createFileIfNotExists(file)
}

func createFileIfNotExists(file string) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) == false {
		return nil
	}

	fileHandle, err := os.Create(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileHandle.Close()

	return nil
}

func (d *Dber) Destroy() error {
	err := os.Remove(d.file)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
