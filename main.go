package main

import "fmt"

func main() {
	cards := []string{"5♤", newCard()}
	cards = append(cards, "6♤")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "A♤"
}
