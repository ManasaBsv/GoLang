package main

import (
	"fmt"
)

func main() {

	i := 2
	switch i {
	case 1:
		a := 42
		fmt.Println(a)
		fmt.Println(&a)

		fmt.Printf("%T\n", a)
		fmt.Printf("%T\n", &a) //Note

		b := &a
		fmt.Println(*b) // b is an int pointer, and the address of a is stored in b, so we can derefence it using * operator, which gives us the value stored at the address
		fmt.Println(*&b)

		*b = 43
		fmt.Println(*b)
	case 2:
		Exec2()
	}

}

type person struct {
	fname   string
	lname   string
	address string
}

func Exec2() {

	p := person{"Rani", "Singh", "Original-Address"}
	fmt.Println("In Exec2 function, before change", p)
	changeMe(&p)
	fmt.Println("In Exec2 function, after change", p)
}

func changeMe(p *person) {

	(*p).address = "Changed-Address" // can also do p.address ,the compiler figures it out for you

}
