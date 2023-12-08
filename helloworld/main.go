package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the home page</h1>\n")

}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(5, 8)
	fmt.Fprintf(w, "<h1>This is the about page</h1>\n")
	fmt.Fprintf(w, "%d + %d = %d\n", 5, 8, sum)
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(9, 3)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
	fmt.Fprintf(w, "%.2f / %.2f = %.2f\n", 9.00, 3.00, result)

}

func divideValues(x, y float32) (float32, error) {
	if x <= 0 || y <= 0 {
		err := errors.New("error: can not devide by zero")
		return 0, err
	}
	return x / y, nil
}

func main() {

	portNum := ":8080"

	http.HandleFunc("/", Home)
	http.HandleFunc("/Divide", Divide)
	http.HandleFunc("/About", About)

	log.Fatal(http.ListenAndServe(portNum, nil))
}
