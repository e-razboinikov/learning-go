package main

import (
	"fmt"
	"net/http"

	"github.com/e-razboinikov/learning-go/3.9/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application has started on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
