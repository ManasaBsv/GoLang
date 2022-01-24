package main

import (
	"fmt"
	"runtime"
)

var y = 10
var x, l int
var n string

// Creating a type
var a int

type numb int

var b numb

//Excercise 2, n1,n2,n3
var n1 int
var n2 string
var n3 bool

//Excercise 3
/*
n2 = "James Bond"
n3 = true ,
Note cannot do this: non-declaration stmt outside func body*/

//Iota
const (
	a1 = 2016 + iota
	a2 = 2016 + iota
	a3 = 2016 + iota
	a4 = 2016 + iota
)

func main() {
	y = 15
	x = 10

	fmt.Println("Inside main, GOOS: ", runtime.GOOS)
	fmt.Println("Inside main, GOARCH: ", runtime.GOARCH)
	fmt.Println("Inside main, x: ", x)
	fmt.Println("Inside main, y: ", y)

	i := 14
	switch i {
	case 1:
		types()
	case 2:
		zero()
	case 3:
		patterns()
	case 4:
		newType()
	case 5:
		excercise1()
	case 6:
		excercise2()
	case 7:
		excercise3AndMore()
	case 8:
		Iota()
	case 9:
		loopExec1()
	case 10:
		loopExec2()
	case 11:
		loopExec3()
	case 12:
		ifExec()
	case 13:
		switchNoExp()
	case 14:
		logicalOp()

	}

}

func types() {
	fmt.Println("Inside types,before assign x: ", x)
	m := 10
	x := 20 // is x:=20 any different from x = 20 here?, yes it is different, so how do I access the global x
	//y = "hey there" , doesn't work,
	y := "hey there" // local y, note :- can also use `` to initialize it with a string, that will also include newline," " wouldn't
	fmt.Println("Inside types, x: ", x)
	fmt.Println("Inside types, y:=(local) : ", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", m)
}

func zero() {
	fmt.Println("Inside zero, l: ", l)
	fmt.Println("Inside zero, s: ", n, "cool")
}

func patterns() {
	fmt.Println("\n\nInside patterns, y=15")
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%v\n", y)
	fmt.Printf("%#v\n", y) // value
	fmt.Printf("%+v\n", y) // for structs, adds field names
	fmt.Printf("%+x\n", y) // go syntax rep. of the value
	fmt.Printf("%#x\n%#X\n", y, y)

	s := fmt.Sprintf("\n\n%#x\t%#X\t", y, y) //turns entire o/p to a string and assigns it to s
	fmt.Println(s)
	fmt.Printf("%T", s)
}

func newType() {
	fmt.Println("\n\nInside newType")
	fmt.Printf("\n\n%T\n", b)
	//a = b, cannot do this, though b is essentially int, but still it is considered as of type 'numb', thus different

	//Conversion:
	a = int(b)     //value of b converted(type numb) converted to type -int
	fmt.Println(a) //This is not casting, but conversion
	c := int(b)
	fmt.Printf("\n%T\n", c) //int , would have been main.numb if c :=b
}

func excercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Printf("\n%v\t%v\t%v\n", x, y, z)
}

func excercise2() {
	fmt.Printf("\n%v\t%v\t%v\n", n1, n2, n3)

}
func excercise3AndMore() {
	n1 = 42
	n2 = "James Bond"
	n3 = true

	/*var num int
	num = 10
	fmt.Println(num) cool, runs*/
	s := fmt.Sprintf("%v\t%v\t%v", n1, n2, n3)
	bs := []byte(n2)
	fmt.Printf("%T", bs)
	fmt.Println("\n", s, bs)
}
func Iota() {
	fmt.Println("\n", a1, a2, a3, a4)
	y := 12
	fmt.Printf("%d\t%b\t%#x\n", y, y, y) //hands on exc(50)
	y = y << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}

func loopExec1() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func loopExec2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i - 64)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func loopExec3() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func ifExec() {

	s := "James Bond"

	if s == "James Bond" {
		fmt.Println("This is Bond")
	} else if s == "Iron Man" {
		fmt.Println("This is Iron Man")
	} else {
		fmt.Println("This is neither Iron Man nor James Bond")
	}
}

func switchNoExp() {

	switch {
	case true:
		fmt.Println("It's true")
	case false:
		fmt.Println("It's false")
	}
}
func logicalOp() {

	if 2 == 2 && 1 == 1 {
		fmt.Println("yAYY")
	}
	if 2 == 2 || 1 == 2 {
		fmt.Println("Yayy again!")
	}
}
