package main

import "fmt"

// global variable
var x int

/*
Data types goes after identifier
variables have nil by default
Constant can be only bool,string or numbers
*/
const y = 2

var name string

func main() {
	printData()

	defer fmt.Println("Bye")
	defer fmt.Println("Good")
	//function scoped variable
	variablewithouttype := "Hello"
	name = "Jhon"
	var message string
	fmt.Println(variablewithouttype)
	fmt.Println(message)
	fmt.Println(name)
	fmt.Println(x)
	fmt.Println("My First Go Program")
	age := 564
	//birthday(&age) //
	fmt.Print((&age))

}

func birthday(pointerAge *int) {

	if *pointerAge > 140 {
		panic("Too old to be true")
	}

	fmt.Printf("The Pointer is %v and the value is %v\n", pointerAge, *pointerAge)

	*pointerAge++

}

func notExported() { //if we use camelCase is private to the package

}
func Exported() { //if we use TitleCase its public and exported to other packages
	id := 4
	price := 56

	priceAsInt := int(price)
	idAsFloat := float32(id)

	fmt.Println(priceAsInt)
	fmt.Println(idAsFloat)

}

// Collections
//Go has Fixed length to the arrays
//Slices :similar  to a dynamic length array.But they are actually chunks of arrays []int

var Countries [10]string

// Go file can have more than one init function
func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"

	fmt.Println("Counties Saved")

	stateTax, _ := calculateTax(1908)
	fmt.Println(stateTax)
}

func init() {

	fmt.Println("A")

}

func init() {

	fmt.Println("B")

}

func calculateTax(price float32) (float32, float32) {

	return price * 0.09, price * 0.01
}

//Error Handling
