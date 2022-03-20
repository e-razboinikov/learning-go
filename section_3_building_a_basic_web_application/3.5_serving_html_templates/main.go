package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

/// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

/// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "abot.page.html")
}

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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Application has started on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
