package main

import "fmt"

func main() {
	cards := []string{"5♤", newCard()}
	cards = append(cards, "6♤")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "A♤"
}
