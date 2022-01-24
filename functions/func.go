package main

import (
	"fmt"
)

func main() {

	i := 4
	switch i {
	case 1:
		foo()
	case 2:
		stringArg("Monica")
	case 3:
		withReturn("Jignes") //Note if return value not accepted, no error.
		s1 := withReturn("Jignes")
		fmt.Println(s1)
	case 4:
		x, y := mreturns("Anisha", "Singh")
		fmt.Println(x, "\n", y)

	}
}

//func (r receiver) identifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("normal function without any parameters/return values")
}

//everything in Go is pass by value(copy created, unlike pass by reference)
func stringArg(s string) { // Note scope of s is just for this function
	fmt.Println("Hello,", s)
}

func withReturn(s string) string {

	return fmt.Sprintf("Hey, %v ", s)
}

//multiple returns
func mreturns(fn, ln string) (string, bool) { //note fn,ln string

	a := fmt.Sprintf(fn, ln, ` Says Hello`)
	b := true
	return a, b
}
