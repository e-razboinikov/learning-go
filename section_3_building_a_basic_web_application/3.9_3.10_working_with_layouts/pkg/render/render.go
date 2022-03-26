package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

/// RenderTemlate - это функция, позволяющая отображать html-шаблоны на экране
func RenderTemplate(writer http.ResponseWriter, templateString string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, ok := templateCache[templateString]
	if !ok {
		log.Fatalf("Template didnt founded, error: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, nil)
	if err != nil {
		fmt.Printf("Error executing template, error: %s", err)
	}

	_, err = buffer.WriteTo(writer)
	if err != nil {
		fmt.Printf("Error writing to buffer, error: %s", err)
	}
}

/// CreateTemplateCache - создает кэш шаблонов типа [map]
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Массив страниц *.page.html
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		fmt.Printf("Error searching templates, error: %s", err)
		return myCache, err
	}

	// Итеррируемся по этому массиву
	for _, page := range pages {
		// Отсекаем путь, кроме *.page.html
		name := filepath.Base(page)

		// Парсим шаблон, добавляя к нему Funcs
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Printf("Error parsing template, error: %s", err)
			return myCache, err
		}

		// Массив разметок *.layout.html
		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			fmt.Printf("Error searching layouts, error: %s", err)
			return myCache, err
		}

		// Парсим шаблоны, сопоставляя их с разметками
		if len(matches) > 0 {
			ts, err := ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				fmt.Printf("Error matching, error: %s", err)
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
