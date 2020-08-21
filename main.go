package main

func main() {
	cards := deck{"5♤", newCard()}
	cards = append(cards, "6♤")
	cards.print()
}

func newCard() string {
	return "A♤"
}
