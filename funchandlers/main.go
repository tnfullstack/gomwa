// Hello Gopher web
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = ":8080"
)

// Home returns the web app home page
func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Home Page.")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Bytes written:", n)

}

// About returns the web app about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 5)
	n, err := fmt.Fprintf(w, "About Page. 2 + 2 is %d", sum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Bytes written:", n)
}

// addVales is a test function for now
func addValues(a, b int) int {
	return a + b
}

// Divide DivideValues return results of to a web page
func Divide(w http.ResponseWriter, r *http.Request) {
	divResult, err := divideValues(9.0, 0)
	if err != nil {
		fmt.Fprintf(w, "Divide page, 9 / 3 is %.2f, error: %s\n", divResult, err)
		return
	}

	fmt.Fprintf(w, "Divide page, 9 / 3 is %.2f\n", divResult)

}

// divideValues calculate two numbers, return result, and error
func divideValues(a, b float64) (float64, error) {
	if a <= 0 || b <= 0 {
		err := errors.New("cannot calculate expression with zero value")
		return 0, err
	} else {
		divRes := a / b
		return divRes, nil
	}
}

// main function is the starting point of the web app.
func main() {

	// HandleFunc listen to web url requests
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting web app on port", webPort)
	// Listen to web server request and serve the pages
	err := http.ListenAndServe(webPort, nil)
	if err != nil {
		log.Println(err)
	}
}
