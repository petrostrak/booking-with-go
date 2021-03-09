package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// Template func will parse the templates
func Template(w http.ResponseWriter, tmpl string) {
	_, err := TemplateTest(w)
	if err != nil {
		return
	}
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	if err = parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// TemplateTest func
func TemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		mathes, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}

		if len(mathes) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
