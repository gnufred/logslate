package report

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
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

//go:embed templates/bots
var botsTemplate string

var templateEmbeds = map[string]string{
	"summary":   summaryTemplate,
	"time":      timeTemplate,
	"threats":   threatsTemplate,
	"attacks":   attacksTemplate,
	"attackers": attackersTemplate,
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

func createTextOutput(r *Report) (bytes.Buffer, error) {
	var out bytes.Buffer

	tmpl := template.New("report")
	tmpl.Funcs(template.FuncMap{
		"HumanTime":     HumanTime,
		"HumanDuration": HumanDuration,
		"AddCommas":     AddCommas,
		"BytesToMB":     BytesToMB,
	})

	templateText, err := compileTemplateText(r)
	if err != nil {
		return out, err
	}

	data, err := prepareData(r)
	if err != nil {
		fmt.Println()
		err := fmt.Errorf("layout specification: %w", err)
		return out, err
	}

	_, err = tmpl.Parse(templateText)
	if err != nil {
		err := fmt.Errorf("template parsing: %w", err)
		return out, err
	}

	err = tmpl.Execute(&out, &data)
	if err != nil {
		err := fmt.Errorf("template execution: %w", err)
		return out, err
	}

	return out, nil
}

func prepareData(r *Report) (*TextReport, error) {
	var err error
	t := NewText(r)
	for _, section := range r.layout {
		switch section.Name {
		case "attackers":
			err = atkrOptions(t, section.Options)
		case "bots":
			err = botOptions(t, section.Options)
		}
	}
	return t, err
}

func atkrOptions(t *TextReport, options Opts) error {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			err = atkrOrder(t, val)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func atkrOrder(r *TextReport, order string) error {
	switch order {
	default:
		return fmt.Errorf("invalid option value \"%s\"", order)
	case "count":
	case "name":
	}

	count := []Count{}
	if order == "count" {
		for ip, atkr := range r.Atkr {
			count = append(count, Count{atkr.Cnt, ip})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	r.AtkrCnt = count
	return nil
}

func botOptions(t *TextReport, options Opts) error {
	var err error
	for opt, val := range options {
		switch opt {
		case "order":
			err = botOrder(t, val)
		default:
			err = fmt.Errorf("invalid option \"%s\"", opt)
		}
	}
	return err
}

func botOrder(t *TextReport, order string) error {
	switch order {
	default:
		return fmt.Errorf("invalid option value \"%s\"", order)
	case "count":
	case "name":
	}

	count := []Count{}
	if order == "count" {
		for name, bot := range t.Bot {
			count = append(count, Count{bot.Cnt, name})
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i].Num > count[j].Num
		})
	}
	t.BotCnt = count
	return nil
}

func compileTemplateText(r *Report) (string, error) {
	text := ""
	for _, section := range r.layout {
		embed, ok := templateEmbeds[section.Name]
		if !ok {
			return "", fmt.Errorf("section \"%s\" does not exist", section.Name)
		}
		text += embed
	}
	return text, nil
}

func writeReport(path string, out []byte) error {
	return os.WriteFile(path, out, 0644)
}
