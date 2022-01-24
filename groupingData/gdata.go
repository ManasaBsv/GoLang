package main

import "fmt"

func main() {

	i := 8
	switch i {
	case 1:
		slicingTheSlice()
	case 2:
		withRange()
	case 3:
		appending()
	case 4:
		ArrayExec()
	case 5:
		SliceExec()
	case 6:
		SliceMake()
	case 7:
		TwoDslice()
	case 8:
		MakeMap()
	}
}

func slicingTheSlice() {

	v := []int{2, 3, 4, 5}

	for i := 0; i < 4; i++ {
		fmt.Print(v[i:i+1], "\t")
	}
}

func withRange() {
	m := []int{2, 3, 4, 5}

	for i, v := range m {
		fmt.Println(i, v)
	}

}
func appending() {
	m := []int{2, 3, 4, 5}

	n := []int{10, 11, 12, 13}

	m = append(m, n...)
	fmt.Println(m)

	//deleting by appending:
	m = append(m[:2], m[5:]...) // deleting 4,5,10
	fmt.Println(m)

}

func ArrayExec() {
	var x [5]int // Can also do x:= [5]int{1,2,3,4,5} , here we've specified the size, i.e. 5
	fmt.Println("Before", x)
	for i := 0; i < 5; i++ {
		x[i] = (i + 1) * 10
	}
	fmt.Println("After", x)

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func SliceExec() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n\n", x) //Note the difference, [5]int, and here []int

	//Slicing a slice
	fmt.Println("\n\nSlicing a slice")
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	//appending
	fmt.Println("\n\nAppending")
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	//Deleting
	fmt.Println("\n\nDeleting")

	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}

func SliceMake() {

	stateNames := []string{`Alabama`, `Arkansas`, `Arizona`, `Alaska`, `California`, `Colorado`, `Connecticut`, `Delaware`,
		`Florida`, `Georgia`, `Hawaii`, `Iowa`, `Illinois`, `Idaho`, `Indiana`, `Kentucky`, `Kansas`, `Loiusiana`, `Maine`, `Maryland`,
		`Massachusetts`, `Minnesota`, `Michigan`, `Missouri`, `Montana`, `Mississippi`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`,
		`North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennyslavania`, `Rhode Island`, `South Carolina`, `South Dakota`,
		`Tenesse`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`}

	fmt.Print(len(stateNames))
	states := make([]string, 49, 60)
	fmt.Println("\nBefore:\n", states)
	for i := 0; i < 49; i++ {
		states[i] = stateNames[i]
	}

	fmt.Println("Before appending, length ", len(states))
	states = append(states, `Wyonming`)
	fmt.Println("After Appending, length", len(states))
	fmt.Println("After:\n", states)

	for i := 0; i < len(states); i++ {
		fmt.Println(" i: ", i, states[i])
	}
}

func TwoDslice() {

	slice1 := []string{"James", "Bond", "Shaken"}
	slice2 := []string{"Miss", "Money"}

	slices := [][]string{slice1, slice2}
	fmt.Println(slices)
	fmt.Println(len(slices))
	fmt.Println(slices[0][1])

	//Rangine over
	fmt.Print("\n\n")
	for i, v := range slices {
		fmt.Println(i)
		for j, v1 := range v {
			fmt.Println("\t", slices[i][j], v1) //can access the value using both the methods, can throw away one(i or v)
		}
	}
}

func MakeMap() {

	m := map[string][]string{
		"Harry":    []string{"Ginny", "Quidditch", "Invisibility Cloack", "Hedwig"},
		"Hermione": []string{"Crookshanks", "Books", "Ron"},
		"Ron":      {"Chocolate Frogs", "Chudley Cannons", "Hermione"}, //note []string not needed.
	}

	fmt.Println(m)
	//Loop Over

	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}

	//Add a record
	m["Neville"] = []string{"Herbology", "Gryffindor", "Parents", "Toad-Trevor"}
	fmt.Printf("\nRecord added:\n\n")
	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}

	//Delete a record
	if v, ok := m["Ron"]; ok {
		delete(m, "Ron")
		fmt.Printf("\nRecord deleted: %v\n", v)

		for key, value := range m {
			fmt.Println(key)
			for _, v := range value {
				fmt.Println("\t", v)
			}
		}
	}

}
