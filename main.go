package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.csv")
	// hand, remainingCards := deal(cards, 3)

	// fmt.Println("My Hand")
	// fmt.Println(hand.toString())

	// fmt.Println("Remaining Deck")
	// fmt.Println(remainingCards.toString())
}
