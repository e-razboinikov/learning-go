package render

import (
	"fmt"
	"net/http"
	"text/template"
)

/// RenderTemlate - это функция, позволяющая отображать html-шаблоны на экране
func RenderTemplate(w http.ResponseWriter, tmpl string) {
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
