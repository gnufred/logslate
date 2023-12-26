package report

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

//go:embed templates/summary
var summaryTemplate string

//go:embed templates/time
var timeTemplate string

//go:embed templates/threats
var threatsTemplate string

//go:embed templates/attacks
var attacksTemplate string

//go:embed templates/attackers
var attackersTemplate string

//go:embed templates/attacklog
var attacklogTemplate string

//go:embed templates/ads
var adsTemplate string

//go:embed templates/bots
var botsTemplate string

var templateEmbeds = map[string]string{
	"summary":   summaryTemplate,
	"time":      timeTemplate,
	"threats":   threatsTemplate,
	"attacks":   attacksTemplate,
	"attackers": attackersTemplate,
	"attacklog": attacklogTemplate,
	"ads":       adsTemplate,
	"bots":      botsTemplate,
}

func (r *Report) Print() {
	switch r.format {
	case "text":
		printText(r)
	case "json":
		printJson(r)
	}
}

func printJson(r *Report) error {
	lerr := log.New(os.Stderr, "", 0)

	out, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		lerr.Println(err)
		return err
	}
	if err := writeReport(r.path, out); err != nil {
		lerr.Println(err)
		return err
	}
	return nil
}

func printText(r *Report) error {
	lerr := log.New(os.Stderr, "", 0)

	out, err := createTextOutput(r)
	if err != nil {
		err := fmt.Errorf("report error: %w", err)
		lerr.Println(err)
		return err
	}

	if r.path == "screen" {
		fmt.Print(out.String())
		return nil
	}

	err = writeReport(r.path, out.Bytes())
	if err != nil {
		lerr.Println(err)
		return err
	}
	return nil
}

func newTemplate() *template.Template {
	tmpl := template.New("report")
	tmpl.Funcs(template.FuncMap{
		"HumanTime":     HumanTime,
		"HumanDuration": HumanDuration,
		"AddCommas":     AddCommas,
		"BytesToMB":     BytesToMB,
	})
	return tmpl
}

func createTextOutput(r *Report) (bytes.Buffer, error) {
	var out bytes.Buffer

	data, err := prepareData(r)
	if err != nil {
		fmt.Println()
		err := fmt.Errorf("layout specification: %w", err)
		return out, err
	}

	for _, section := range r.layout {
		templateText, ok := templateEmbeds[section.Name]
		if !ok {
			return out, fmt.Errorf("section \"%s\" does not exist", section.Name)
		}

		tmpl := newTemplate()

		_, err = tmpl.Parse(templateText)
		if err != nil {
			err := fmt.Errorf("template parsing: %w", err)
			return out, err
		}

		buf := bytes.Buffer{}
		err = tmpl.Execute(&buf, &data)
		if err != nil {
			err := fmt.Errorf("template execution: %w", err)
			return out, err
		}

		bytes := buf.Bytes()
		if len(bytes) > 1 {
			out.WriteString(header(section.Name))
			out.Write(bytes)
		}
	}

	return out, nil
}

func header(name string) string {
	switch name {
	case "attacklog":
		name = "Attack Log"
	default:
		name = strings.ToUpper(name[:1]) + name[1:]
	}
	return fmt.Sprintf("=== [[ %s ]] ===", name)
}

func prepareData(r *Report) (*TextReport, error) {
	var err error
	t := NewText(r)
	for _, section := range r.layout {
		switch section.Name {
		case "attackers":
			fd := filterData{List: t.Attackers}
			fd, err = execOptions(fd, section.Options)
			t.Attackers = fd.List
			t.AttackersFiltered = fd.Filtered
			t.AttackersCount = fd.Cnt
		case "ads":
			fd := filterData{List: t.Ads}
			fd, err = execOptions(fd, section.Options)
			t.Ads = fd.List
			t.AdsFiltered = fd.Filtered
			t.AdsCount = fd.Cnt
		case "bots":
			fd := filterData{List: t.Bots}
			fd, err = execOptions(fd, section.Options)
			t.Bots = fd.List
			t.BotsFiltered = fd.Filtered
			t.BotsCount = fd.Cnt
		}
	}
	return t, err
}

func writeReport(path string, out []byte) error {
	return os.WriteFile(path, out, 0644)
}
