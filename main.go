package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	fmt.Println("My Hand")
	fmt.Println(hand.toString())

	fmt.Println("Remaining Deck")
	fmt.Println(remainingCards.toString())
}
