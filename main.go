package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.saveToFile("cards.csv")
	// cards := newDeckFromFile("cards.csv")
	// cards.print()
	// hand, remainingCards := deal(cards, 3)

	// fmt.Println("My Hand")
	// fmt.Println(hand.toString())

	// fmt.Println("Remaining Deck")
	// fmt.Println(remainingCards.toString())
}
