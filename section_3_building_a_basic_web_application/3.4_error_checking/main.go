package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0

	result, err := dividion(x, y)

	if err != nil {
		fmt.Fprintf(w, fmt.Sprint(err))
		return
	}

	fmt.Fprintf(w, "%f divided by %f equals to %f", x, y, result)
}

func dividion(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}

	return x / y, nil
}

func main() {
	http.HandleFunc("/", Divide)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
