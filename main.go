package main

import "fmt"

// custom type struct
// type contactInfo struct {
// 	email   string
// 	zipCode int
// }

// type person struct {
// 	firstName string
// 	lastName  string

// 	// Embedding Structs
// 	// => Embedding one struct into another
// 	contact contactInfo
// }

// func main() {
// 	// creating an object using struct
// 	// alex will be of type person
// 	// alex := person { "Alex", "Anderson" }

// 	// If we want to change the order of assignment of values
// 	// i.e., first word will be lastname & second word will be firstName
// 	// alex := person {firstName : "Alex", lastName : "Anderson"}

// 	// fmt.Println(alex)

// 	// Method-II of creating object
// 	// var alex person // Here, firstName and lastName will be creating but, they will be empty strings
// 	// fmt.Println(alex)

// 	// Type  ZeroValue(default)
// 	// string   ""
// 	// int      0
// 	// float    0
// 	// bool     false

// 	// It prints keys, values of person object
// 	// fmt.Printf("%+v", alex)

// 	// // updating values of object
// 	// alex.firstName = "Alex"
// 	// alex.lastName = "Anderson"

// 	// fmt.Printf("%+v", alex)

// 	// Embedding Structs
// 	// => Embedding one struct into another
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 96879,
// 		},
// 	}

// 	// jim.updateName("Jimmy")
// 	// jim.print()

// 	// Structs with pointers
// 	// jimPointer := &jim
// 	// jimPointer.updateName("jimmy")
// 	// jim.print()

// 	jim.updateName("jimmy")
// 	jim.print()
// }

// // Receiver functions
// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// // func (p person) updateName(newFirstName string) {
// // 	p.firstName = newFirstName
// // }

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

// slice works differently from struct
// func main() {
// 	mySlice := []string{"Hi", "There", "How", "are", "you"}

// 	updateSlice(mySlice) // slice gets updated

// 	fmt.Println(mySlice) // o/p : Bye There How are you
// }

// func updateSlice(s []string) {
// 	s[0] = "Bye"
// }

// IMPP : slices, maps, channels, pointers, functions => are reference types
// => we don't need to worry about pointers with these
// => Go automatically understands them

// Value Types => int, float, string, bool, structs  => we need to use pointers to change these things

// MAPS
// map => Go
// hash => Ruby
// Object => JavaScript
// Dict => python

// map datastructure => stores key->value pairs

// Iterating over the map
func printMap(mp map[string]string) {
	for color, hex := range mp {
		fmt.Println(color, hex)
	}
}

func main() {
	// creating a map
	// key of type string & value of type string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// fmt.Println(colors)

	colors := make(map[string]string) // built in function creates empty map

	colors["white"] = "#ffffff" // inserting an entry
	colors["red"] = "#ff0000"

	// deleting an entry
	// delete(colors, "white")

	// fmt.Println(colors)

	// Iterating over the map
	printMap(colors)
}

// Difference b/w struct and maps

// Map
// 1) All keys must be of same type
// 2) All values must be of same type
// 3) Keys are indexed - we can iterate over them
// 4) Reference Type

// struct
// 1) Values can be of different type
// 2) keys doesn't support indexing
// 3) Value Type
