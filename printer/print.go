package printer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
)

type Printer struct {
	Path        string
	Format      args.LogFormat
	File        *os.File
	NewLine     string
	Stringifier func(any) (string, error)
}

func (p *Printer) OpenFile() error {
	file, err := os.OpenFile(
		p.Path,
		os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Printf("Error: Can't open: %s\n", p.Path)
		return err
	}
	p.File = file

	return nil
}

func (p *Printer) CloseFile() {
	if p.File != nil && p.Format != args.SCREEN {
		p.File.Close()
	}
}

func Error(err error) {
	lerr := log.New(os.Stderr, "", 0)
	perr := fmt.Errorf("printer error: %w", err)
	lerr.Println(perr)
}

func (r *Printer) Print(data any) error {
	// Both need stringification
	if r.Format == args.SCREEN || r.Format == args.TEXT {
		if r.Stringifier == nil {
			err := errors.New("no stringifier set")
			Error(err)
			return err
		}

		s, err := r.Stringifier(data)
		if err != nil {
			Error(err)
			return err
		}

		switch r.Format {
		case args.SCREEN:
			return r.PrintScreen(s)
		case args.TEXT:
			return r.PrintText(s)
		}
	}

	if r.Format == args.JSON {
		return r.PrintJson(data)
	}

	return errors.New("no format specified")
}

func (p *Printer) PrintScreen(s string) error {
	_, err := fmt.Print(s)
	return err
}

func (p *Printer) PrintText(s string) error {
	return p.WriteString(s)
}

func (p *Printer) PrintJson(data any) error {
	out, err := json.Marshal(data)
	if err != nil {
		Error(err)
		return err
	}
	return p.WriteString(string(out) + p.NewLine)
}

func (p *Printer) WriteString(out string) error {
	_, err := p.File.WriteString(out)
	if err != nil {
		Error(err)
	}
	return err
}
