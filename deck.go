package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♤", "♢", "♡", "♧"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	bs := []byte(d.toString())
	return ioutil.WriteFile(filename, bs, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Errors", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	r := newRand()
	max := len(d) - 1
	for i := range d {
		d.swap(i, r.Intn(max))
	}
}

func (d deck) swap(oldI int, newI int) {
	d[oldI], d[newI] = d[newI], d[oldI]
}

func newRand() *rand.Rand {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	return rand.New(source)
}
