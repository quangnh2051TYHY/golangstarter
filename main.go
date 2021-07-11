package main

import (
	"fmt"
	"io"
	"os"
)

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
	// listCard = append(listCard, "Quang")
	// fmt.Println(listCard)

	// for i, cardElement := range listCard {
	// 	fmt.Println(i, cardElement)
	// }
	listCard = getNewDeck()
	// listCard.printFromDeckClass()
	fmt.Println("Tostring" + listCard.deckToString())
	left, right := deal(listCard, 2)
	left.printFromDeckClass()
	right.printFromDeckClass()
	right.saveStringToFile("abc.text")
	eb := engLishBotChat{name: "ENglish bot"}
	eb.setName("a")
	fmt.Println(eb.name)
	vb := vietNamBotChat{}
	playGreet(eb)
	fmt.Println(vb.getGreet("vietn"))
	fmt.Println(eb.getVersion())

	google := webInfo{
		url: "http://google.com",
	}
	res, err := getRequest(google)
	//under code is the same
	// byteSlide := make([]byte, 15000)
	// //read resBody to byteSlide
	// res.Body.Read(byteSlide)
	// fmt.Println(string(byteSlide))
	fmt.Println(err)
	io.Copy(os.Stdout, res.Body)
	//custom an Writer Interface like upward code
	cw := CustomWriter{}
	io.Copy(cw, res.Body)
}

type CustomWriter struct {
}

func (CustomWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Length of byte slide", len(bs))
	return len(bs), err
}
func printText() {
	fmt.Println("this is internal call")
}

func testReturn() string {
	return "has type of String"
}
