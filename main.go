package main

import "fmt"

func main() {
	cards := []string{"5♤", newCard()}
	cards = append(cards, "6♤")

	fmt.Println(cards)
}

func newCard() string {
	return "A♤"
}
