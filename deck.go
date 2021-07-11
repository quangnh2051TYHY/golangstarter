package main

import (
	"fmt"
	"io/ioutil"

	// "os"
	"strings"
)

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
			fmt.Println(i + j)
			firstDecksList = append(firstDecksList, firstDeck+secondDeck)
		}
	}
	return firstDecksList
}

func deal(d deck, index int) (deck, deck) {
	return d[:index], d[index:]
}

//code quang vieets
// func deckToString(d deck) string {
// 	stringReturn := ""
// 	for i, deckElement := range d {
// 		fmt.Print(i)
// 		stringReturn += deckElement
// 	}
// 	return stringReturn
// }

//code giao vien
//join
func (d deck) deckToString() string {
	//join what []string, separate by
	return strings.Join([]string(d), "")
}
func (d deck) saveStringToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.deckToString()), 0666)
}

// func (deck) printDeck() string{
// 	return
// }

// func readDeckFromFile(fileName string) deck {
// 	byteSlide, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// }
