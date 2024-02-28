package report

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"strings"
)

//go:embed templates/parsed
var parsedTemplate string

//go:embed templates/time
var timeTemplate string

//go:embed templates/scanned
var scannedTemplate string

//go:embed templates/threats
var threatsTemplate string

//go:embed templates/attacks
var attacksTemplate string

//go:embed templates/attackers
var attackersTemplate string

//go:embed templates/ads
var adsTemplate string

//go:embed templates/bots
var botsTemplate string

var templateEmbeds = map[string]string{
	"parsed":    parsedTemplate,
	"time":      timeTemplate,
	"scanned":   scannedTemplate,
	"threats":   threatsTemplate,
	"attacks":   attacksTemplate,
	"attackers": attackersTemplate,
	"ads":       adsTemplate,
	"bots":      botsTemplate,
}

var Stringify = func(data any) (string, error) {

	r, ok := data.(Report)
	if !ok {
		return "", errors.New("invalid Report struct data")
	}

	tr, err := prepareData(&r)
	if err != nil {
		return "", fmt.Errorf("layout specification: %w", err)
	}

	out, err := createTextOutput(tr)
	if err != nil {
		return "", fmt.Errorf("report error: %w", err)
	}

	return out.String(), nil
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

func header(name string) string {
	name = strings.ToUpper(name[:1]) + name[1:]
	return fmt.Sprintf("=== [[ %s ]] ===", name)
}

func prepareData(r *Report) (*TextReport, error) {
	var err error
	t := NewText(r)
	for _, section := range r.Layout {
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

func createTextOutput(r *TextReport) (bytes.Buffer, error) {
	var out bytes.Buffer

	for _, section := range r.Layout {
		templateText, ok := templateEmbeds[section.Name]
		if !ok {
			return out, fmt.Errorf("section \"%s\" does not exist", section.Name)
		}

		tmpl := newTemplate()

		_, err := tmpl.Parse(templateText)
		if err != nil {
			err := fmt.Errorf("template parsing: %w", err)
			return out, err
		}

		buf := bytes.Buffer{}
		err = tmpl.Execute(&buf, &r)
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
