package main

import "fmt"

func main() {
	// var card_dragon string = "";
	// card_dragon = fmt.Scan()
	var card string = "Yugioh card"
	card += " " + testReturn()
	// card := "acv"
	fmt.Println("This is main " + card)

	printText()

	//Array in go
	// var listCard []string = []string{"a", "b"}
	var listCard deck = deck{"a", "b"}
	//append element to list
	listCard = append(listCard, "Quang")
	fmt.Println(listCard)

	// for i, cardElement := range listCard {
	// 	fmt.Println(i, cardElement)
	// }
	listCard = getNewDeck()
	listCard.printFromDeckClass()
}

func printText() {
	fmt.Println("this is internal call")
}

func testReturn() string {
	return "has type of String"
}
