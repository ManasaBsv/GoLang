package main

import "fmt"

type person struct {
	first string
	last  string
}

type personExec struct {
	firstname string
	lastname  string
	fav_ice   []string
}

func main() {

	i := 5
	switch i {
	case 1:
		NormalStruc()
	case 2:
		EmbeddedStruc()
	case 3:
		structExec1()
	case 4:
		structExec2()
	case 5:
		structAnonymous()
	}

}

func NormalStruc() {

	p1 := person{first: "James", last: "Bond"}
	p2 := person{
		first: "William",
	} // Note ',' needed if } not on the same line

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last)

	fmt.Printf("%T\t%+v", p1, p1) //Note +v
}

func EmbeddedStruc() {

	type secretAgent struct {
		person // this value will be filled by composite literal
		ltk    bool
	}

	sa1 := secretAgent{
		person: person{
			first: "Lily", last: "Jones",
		},
		ltk: true,
	}
	fmt.Println(sa1, sa1.first, sa1.last, sa1.ltk) // so can directly acess the parent structure variables as well

	//Lets try some things out
	type sA struct {
		person // this value will be filled by composite literal
		last   string
	}

	s := sA{
		person: person{
			first: "Lily", last: "Jones",
		},
		last: "Trying something out",
	}

	fmt.Println(s, s.first, s.last) // See overriding. Can do person.last (in case of name collisions),
	//Note that inner type is promoted to outer type(inheritance)
	type sAFinal struct {
		personDetails person // this value will be filled by composite literal
		last          string
	}

	//Can also do like:
	s1 := sAFinal{
		personDetails: person{
			first: "Amelia",
			last:  "Jones",
		},
		last: "Worked",
	}
	fmt.Println(s1, s1.personDetails.first, s1.personDetails.last, s1.last) // See works.

}

func structExec1() {

	pe1, pe2 := personExec{
		firstname: "Rani",
		lastname:  "Singh",
		fav_ice:   []string{"Chocolate", "Strawberry", "Pista"},
	}, personExec{
		firstname: "Karishma",
		lastname:  "Shembde",
		fav_ice:   []string{"Chocolate", "Cassata"},
	}

	pes := []personExec{pe1, pe2}
	fmt.Println(pe1, "\n", pe2)
	//Ranging over

	fmt.Printf("\n\nRanging over structrure Array\n\n")
	mapPerson := map[string]personExec{}

	for i, v := range pes {

		fmt.Print(i, "\n\tName: ", v.firstname, " ", v.lastname, "\n\tIceCreams: ")
		for _, m := range v.fav_ice {
			fmt.Print(m, "\t")
		}
		fmt.Println()

		mapPerson[v.lastname] = v

	}

	fmt.Printf("\n\nRanging over Map\n\n")
	for key, value := range mapPerson {

		fmt.Println(key, "\n\t", value.firstname, value.lastname)
		fmt.Printf("\t")
		for _, i := range value.fav_ice {
			fmt.Print(i, "\t")
		}
		fmt.Println()
	}

}

func structExec2() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: true,
	}
	fmt.Println(t1, "\n", s1)
	fmt.Println("\nTruck: Doors", t1.doors, "\tColor: ", t1.color, "\tFourWheeler: ", t1.fourWheel)

	fmt.Println("\nSedan: Doors", s1.doors, "\tColor: ", s1.color, "\tLuxury: ", s1.luxury)
}

func structAnonymous() {

	p1 := struct {
		name        string
		credentials map[string]string
		websites    []string
	}{
		name: "Priya Sharma",
		credentials: map[string]string{
			"priyaS@gmail": "gmailpWD",
			"priyaS@yahoo": "yahoopWD",
		},
		websites: []string{"gmail", "yahoomail"},
	}

	fmt.Println("Name: ", p1.name)
	fmt.Print("Credentials: \n")
	for key, value := range p1.credentials {
		fmt.Println("\t", key, ": ", value)
	}
	fmt.Print("Webistes: \n\t")
	for _, value := range p1.websites {
		fmt.Print(value, "\t")
	}
}
