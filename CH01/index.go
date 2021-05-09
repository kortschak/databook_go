//go:generate go run index.go *SEC*.md

package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

func main() {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	files, err := filepath.Glob(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(files)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("page").Funcs(map[string]interface{}{
		"basename": func(s string) string {
			return s[:len(s)-len(filepath.Ext(s))]
		},
	}).Parse(`# {{.Title}}
{{range .Sections}}
## [{{basename .}}]({{.}})
{{end}}`))
	tmpl.Execute(f, page{Title: filepath.Base(dir), Sections: files})
	if err != nil {
		log.Fatal(err)
	}

}

type page struct {
	Title    string
	Sections []string
}
