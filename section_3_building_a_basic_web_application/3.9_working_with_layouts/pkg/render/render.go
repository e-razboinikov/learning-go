package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

/// RenderTemlate - это функция, позволяющая отображать html-шаблоны на экране
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemlateTest(w)
	if err != nil {
		fmt.Println(err)
		return
	}

	/// Парсинг шаблона
	parsedTemplate, err := template.ParseFiles("../../templates/" + tmpl)
	/// Обработка ошибок
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

	/// Отображение шаблона
	err = parsedTemplate.Execute(w, nil)
	/// Обработка ошибок
	if err != nil {
		fmt.Println("error executing template: ", err)
		return
	}
}

func RenderTemlateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		fmt.Println(err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println("Page is currently", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println(err)
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
