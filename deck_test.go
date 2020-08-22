package main

import (
	"os"
	"testing"
)

func TestNewDeck52Cards(t *testing.T) {
	d := newDeck()
	expected := 52
	actual := len(d)
	if expected != actual {
		t.Errorf("Expected %v cards, got %v", expected, actual)
	}
}

func TestNewDeckFormat(t *testing.T) {
	d := newDeck()
	expected := "A♤"
	actual := d[0]
	if expected != actual {
		t.Errorf("Expected %v format, got %v", expected, actual)
	}
}

func TestNewDeckCheckValues(t *testing.T) {
	d := newDeck()
	expected := "A♤"
	actual := d[0]
	if expected != actual {
		t.Errorf("Expected %v at first, got %v", expected, actual)
	}

	expected = "2♤"
	actual = d[1]
	if expected != actual {
		t.Errorf("Expected %v at second, got %v", expected, actual)
	}

	expected = "K♧"
	actual = d[51]
	if expected != actual {
		t.Errorf("Expected %v at last, got %v", expected, actual)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck.test")

	d := newDeck()
	d.saveToFile("_deck.test")

	loadedDeck := newDeckFromFile("_deck.test")

	expected := 52
	actual := len(loadedDeck)
	if expected != actual {
		t.Errorf("Expected %v cards, got %v", expected, actual)
	}
}
