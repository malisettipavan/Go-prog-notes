package main

import (
	"fmt"
	"strings"
)

// types and functions of deck Object

// Custom Type Declarations  => Using for Object Oriented programming in Go

// Create a new Type of 'deck' which is a slice of strings
type deck []string

// newDeck function -> it doesnot require receiver as we are just creating a new deck
func newDeck() deck {
	// returns a list of playing cards deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ => is used to avoid compiler unused errors i.e., we should not use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// (cards deck) => Receiver argument => cards is the reference to the actual type deck
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// MULTIPLE RETURN VALUES

// function
func deal(d deck, handSize int) (deck, deck) {
	// seperating the deck into two parts
	return d[:handSize], d[handSize:]
}

// convert deck(list of cards) to string => Join cards to make them to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
