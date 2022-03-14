package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// Home - обработчик перехода на домашнюю страницу
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page!")
}

// About - это обработчик перехода на страницу about
func About(w http.ResponseWriter, r *http.Request) {
	x := 14
	y := 29
	sum := addition(x, y)
	fmt.Fprintf(w, "%d + %d = %d", x, y, sum)
}

func addition(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/about", About)
	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
