package main

import "fmt"

//create a new type of 'Deck'
//that is a slice of strings

type deck []string

//coi d nhỏ như this trong java
func (d deck) printFromDeckClass() {
	for i, cardElement := range d {
		fmt.Println(i, cardElement)
	}
}

func getNewDeck() deck {
	// var newDecksList deck = deck{"newA", "newB"}
	//init deckList empty
	var firstDecksList deck = deck{}
	firstDecksList = []string{"firstA", "firstB"}
	secondDeckList := []string{"deck2", "deck2.1"}

	for i, firstDeck := range firstDecksList {
		for j, secondDeck := range secondDeckList {
			fmt.Println(i+j, firstDeck+" and "+" : "+secondDeck)
		}
	}
	return firstDecksList
}
