package main

import (
	"fmt"

	"strings"
)

// Create a new type of "deck" which is a slice of strings

type deck []string


func newDeck() deck {
	cards := deck{}
	
	cardSuits:= []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards


}


func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}


func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]




}



func (d deck) toString() string {

	return strings.Join([]string(d), ",")
	 

}