package report

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
)

//go:embed main.txt
var mainTemplate string

func (r *Report) Print() {
	switch r.Frmt {
	case "text":
		printText(r)
	case "json":
		printJson(r)
	}
}

func printJson(r *Report) {
	lerr := log.New(os.Stderr, "", 0)

	out, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		lerr.Println(err)
	}
	if err := writeReport(r.Path, out); err != nil {
		lerr.Println(err)
	}
}

func printText(r *Report) {
	lerr := log.New(os.Stderr, "", 0)

	out, err := createTextOutput(r)
	if err != nil {
		lerr.Println(err)
	}
	if r.Path == "screen" {
		fmt.Print(out.String())
	} else {
		if err := writeReport(r.Path, out.Bytes()); err != nil {
			lerr.Println(err)
		}
	}
}

func createTextOutput(r *Report) (bytes.Buffer, error) {
	var out bytes.Buffer
	tmpl := template.New("report")
	tmpl.Funcs(template.FuncMap{
		"HumanTime":     HumanTime,
		"HumanDuration": HumanDuration,
		"AddCommas":     AddCommas,
		"BytesToMB":     BytesToMB,
	})
	_, err := tmpl.Parse(mainTemplate)
	if err != nil {
		return out, err
	}
	err = tmpl.Execute(&out, &r)
	if err != nil {
		return out, err
	}
	return out, nil
}

func writeReport(path string, out []byte) error {
	return os.WriteFile(path, out, 0644)
}
