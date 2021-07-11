package main

import "fmt"

type Ibot interface {
	getGreet(s string) string
	getVersion() string
}
type engLishBotChat struct {
	name string
}
type vietNamBotChat struct{}

func playGreet(b Ibot) {
	fmt.Println(b.getGreet(" is real"))
}

func getOut(b Ibot) {
	fmt.Println(b.getVersion())
}

func (engLishBotChat) getGreet(s string) string {
	return "ENglish" + s
}

func (vietNamBotChat) getGreet(s string) string {
	return "VietNam" + s
}

func (engLishBotChat) getVersion() string {
	return "verson 1"
}

func (eb engLishBotChat) setName(name string) {
	eb.name = "newname and not pointer"
}
