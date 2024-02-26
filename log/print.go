package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitlab.com/gnufred/logslate/args"
)

type ReportTL struct {
	path   string
	format args.LogFormat
	file   *os.File
	Entry  ThreatLog
}

func New(path string, format args.LogFormat) *ReportTL {
	r := ReportTL{}
	r.path = path
	r.format = format

	if r.format != args.SCREEN {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error: Can't open: %s\n", r.path)
		}
		r.file = file
	}

	return &r
}

func (r *ReportTL) CloseFile() {
	if r.file != nil && r.format != args.SCREEN {
		r.file.Close()
	}
}

func (r *ReportTL) Print() {
	switch r.format {
	case args.SCREEN:
		printScreen(r)
	case args.TEXT:
		printText(r)
	case args.JSON:
		printJson(r)
	}
}

func printScreen(r *ReportTL) {
	fmt.Print(compileEntryString(r.Entry) + "\n")
}

func printText(r *ReportTL) error {
	lerr := log.New(os.Stderr, "", 0)
	out := compileEntryString(r.Entry)
	err := writeReport(r.file, r.path, out)
	if err != nil {
		lerr.Println(err)
		return err
	}
	return nil
}

func printJson(r *ReportTL) error {
	lerr := log.New(os.Stderr, "", 0)

	out, err := json.Marshal(r.Entry)
	if err != nil {
		lerr.Println(err)
		return err
	}
	if err := writeReport(r.file, r.path, string(out)); err != nil {
		lerr.Println(err)
		return err
	}
	return nil
}

func compileEntryString(t ThreatLog) string {
	var s string
	sep := " "
	q := "\""

	s = s + q + t.Time.String() + q

	if t.Ip != nil {
		s = s + sep + t.Ip.String()
	}

	if t.IsBot() {
		s = s + sep + "bot:" + t.Bot
	}

	if t.IsAd() {
		s = s + sep + "adv:" + t.Ad
	}

	if t.IsAttack() {
		s = s + sep + "att:" + t.Attack.Name + sep + url(t.Path, t.Query)
	}

	return s
}

func writeReport(file *os.File, path string, out string) error {
	_, err := file.WriteString(out + "\n")
	return err
}

func url(p, q string) string {
	if p != "" && q != "" {
		return p + "?" + q
	}
	return p + q
}
