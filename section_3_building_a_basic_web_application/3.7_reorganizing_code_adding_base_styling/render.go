package main

import (
	"fmt"
	"net/http"
	"text/template"
)

/// renderTemlate - это функция, позволяющая отображать html-шаблоны на экране
func renderTemplate(w http.ResponseWriter, tmpl string) {
	/// Парсинг шаблона
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
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
