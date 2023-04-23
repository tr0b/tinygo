package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	edl := 16              // expected deck length for test
	dl := len(d)           // deck length
	fc := d[0]             // first card of deck
	lc := d[len(d)-1]      // Last card of deck
	efc := "Ace of Spades" // Expected First card of deck
	elc := "Four of Clubs" // Expected Last card of deck

	if dl != edl {
		t.Errorf("Expected deck length of %d, but got %d", edl, dl)
	}

	if fc != efc {
		t.Errorf("Expected first card of %v, but got %v", efc, fc)
	}

	if lc != elc {
		t.Errorf("Expected last card of %v, but got %v", elc, lc)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	expectedDeckLength := 16 // expected deck length for test

	loadedDeck := newDeckFromFile("_decktesting")
	loadedDeckLength := len(loadedDeck) // deck length

	if loadedDeckLength != expectedDeckLength {
		t.Errorf("Expected %v cards in deck, got %v", expectedDeckLength, loadedDeckLength)
	}

	os.Remove("_decktesting")
}
