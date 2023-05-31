package data

import "fmt"

// Collections are not objects(nothing is an object actually) so we user gloabal functions to work with
// them such.such as len() or cap
var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"
	//qty := len(Countries)
	fmt.Print("Countries Saved")

	Prices := [2]int{100.200}
	//Slices
	name := []string{"Mary", "Jhon"}
	names := append(names, "Carol")
	println(len(names)) //length of the name array

	//Map
	wellKnownPorts := map[string]int{"http": 80, "https": 443}

}

//Arrays

//Functions types

func save() {

}

func save_with_paramter(text string) {

}

func add(a int, b int) int {

	return a + b

}

func addAndSubstract(a int, b int) (int, int) {

	return a + b, a - b
}
