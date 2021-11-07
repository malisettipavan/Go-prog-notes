package main

import "fmt"

// If there is package main in go file => we need to declare main() function
// => So, that if we build this go file => we get an executable file
// => And when we run the executable file => main() function gets called

// Basic go types
// 1) bool
// 2) string
// 3) int
// 4) float64

// While writing Global variables
// => we should not use automatic type declaration => name := "pavan" => GIVES ERROR
// => we need to mention variable type on our own
// var name string = "pavan"

// func main() {
// 	// variable car of type string
// 	var card string = "Ace of Spades"

// 	// Assigns variable type automatically according to the assigning value
// 	card2 := "Ace of hearts" // we need to use  " : " for the first time while declaring the variable

// 	// Since, card2 is already declared => we don't have to use  " : "
// 	card2 = "pavan"

// 	fmt.Println(card)
// 	fmt.Println(card2)
// 	fmt.Println(name)
// }

// FUNCTIONS & RETURN TYPES

// func main() {

// 	card := newCard()
// 	fmt.Println(card)

// 	temp := newCardNum()
// 	fmt.Println(temp)

// }

// // Telling the function to return string
// func newCard() string {
// 	return "Five of Diamonds"
// }

// func newCardNum() int {
// 	return 5
// }

// IMPPP : Files of the same Package donot need to be imported
// => we can call a function of file1 in file2

// ARRAYS & SLICES

// In Go, we have two types of data structures to handle lists

// Array => Fixed length list of things
// Slice => Array that can grow (or) shrink

// func main() {
// 	// creating slice
// 	cards := []string{} // we can add elements inside {}
// 	// eg : cards := []string{"pavan", "kalyan"}

// 	// Appending a new element to the slice
// 	cards = append(cards, "shiva")

// 	fmt.Println(cards)

// 	// Iterate over a cards Slice
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

// Object Oriented Programming

func main() {
	cards := newDeck()

	// calling print() function
	// cards.print()

	// deal function returns two values
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	// Slice Range
	// Eg : fruits := []string{"apple", "banana", "grape", "orange"}

	// fruits[startingIndex : EndingIndex] (endingIndex not included)
	// fruits[ : endingIndex] => it automatically takes startingIndex = 0
	// fruits[startingIndex : ] => it automatically takes endingIndex = size of slice

	// fruits := []string{"apple", "banana", "grape", "orange"}
	// fmt.Println(fruits[1:3])

	// fmt.Println(cards)

	// getName("shiva")

	fmt.Println(cards.toString())
}

func getName(name string) {
	fmt.Println(name)
}

func saveToFile() {
	
}
