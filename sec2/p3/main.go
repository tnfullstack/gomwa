// More on variables
package main

import "log"

var s = "June"

func main() {
	var s2 = "May"

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	say, err := saySomething("May and June are the best time for a cross country driving trip...")
	if err != nil {
		log.Println("Something went wrong?")
	}
	log.Println(s)
	log.Printf("String returned from saySomething Func is:\n %s\n", say)
}

func saySomething(s string) (string, error) {
	return s, nil
}
