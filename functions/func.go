package main

import (
	"fmt"
	"math"
)

func main() {

	i := 15
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
	case 5:
		variadic(1, 2, 3, 4, 5, 6)
		x1 := []int{2, 4, 6, 8}
		variadic(x1...) // x1 unfurling and passed to a variadic parameter, here same underlying array used by x, the parameter as well, no new slice created.
		fmt.Println("No arguments")
		variadic()
	case 6:
		deferred()
	case 7:
		callback()
	case 8:
		firstExec()
	case 9:
		secondExec()
	case 10:
		fourthExec()
	case 11:
		fifthExec()
	case 12:
		sixthExec()
	case 13:
		eightExec()
	case 14:
		ninthExec()
	case 15:
		tenthExec()
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

	a := fmt.Sprintf("%v %v Says hello", fn, ln)
	b := true
	return a, b
}

func variadic(x ...int) { // x... variadic, so can pass in 0 arguments or more
	// parameters -into a slice named x of type int  ( x is of type ...T)
	//Also x ...int(the variadic parameter) should be the final parameter : (x ...int, s string), won't work
	fmt.Println(x, len(x), cap(x)) // in case of x being nil(if no args passed), slice created but underlying array didn't get created
	fmt.Printf("%T", x)

}

//defer -> the function 'deferred' will run after the function where it's defined, ends

func deferred() {

	{
		defer fooD()
	}
	bar()
}
func fooD() {
	fmt.Println("FooD")
}
func bar() {
	fmt.Println("Bar")
}

//Methods
//Interfaces - helps us do polymorphism
//callback
func callback() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	total := sum(nums...)
	fmt.Println("Total of all numbers: ", total)

	totalEven := totalEven(sum, nums...)
	fmt.Println("Total of all even numbers: ", totalEven)
}

func sum(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func totalEven(f func(x ...int) int, n ...int) int {

	var even_slice []int
	for i, v := range n {
		if i%2 == 0 {
			even_slice = append(even_slice, v)
		}
	}

	return f(even_slice...)
}

//firstExec
func firstExec() {
	fmt.Println(foo1())
	fmt.Println(bar1())
}
func foo1() int {
	return 4
}
func bar1() (int, string) {
	return 4, "Four"
}

//secondExec

func secondExec() {

	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Called foo2, sum: ", foo2(sl...))
	fmt.Println("Called bar2, sum: ", bar2(sl))

}
func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar2(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//fourthExec

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Name: ", p.first, p.last, " Age: ", p.age)
}
func fourthExec() {
	p1 := person{
		"Dr.",
		"Strange",
		45,
	}
	p1.speak()
}

//fifthExec
type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(sh shape) {

	switch sh.(type) {
	case circle:
		fmt.Println("Circle: Radius: ", sh.(circle).radius, " and the area is: ", sh.area())

	case square:
		fmt.Println("Square: Side: ", sh.(square).side, " and the area is: ", sh.area())
	}
}

func fifthExec() {
	c1 := circle{5}
	s1 := square{5}
	info(c1)
	info(s1)
}

//sixthExec

func sixthExec() {

	func(n int) {
		fmt.Println("Anonymous func number", n)
	}(1)

	var f2 func(int) int //Note
	f1 := func(x int) int {
		sum := 0
		for ; x > 0; x-- {
			sum += x
		}
		return sum
	}
	f2 = f1 //Note the assignment
	fmt.Println("Sum from 1-4:", f1(4))
	fmt.Println(f2(4))

}

//eightExec
func eightExec() {
	i := returnsFunc()

	fmt.Println(i())
}
func returnsFunc() func() int {

	return func() int {
		return 50
	}
}

//ninthExec
func ninthExec() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Even sum:", eosum(sum1, "even", n...))
	fmt.Println("Odd sum:", eosum(sum1, "odd", n...))
}
func sum1(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}
func eosum(f func(...int) int, s string, n ...int) int {
	sl := []int{}
	switch s {
	case "even":
		for i, v := range n {
			if i%2 == 0 {
				sl = append(sl, v)
			}
		}
	case "odd":
		for i, v := range n {
			if i%2 == 1 {
				sl = append(sl, v)
			}
		}

	}
	return f(sl...)
}

//tenthExec - closure
func tenthExec() {
	a := returnClosure()
	b := returnClosure()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
}

func returnClosure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
