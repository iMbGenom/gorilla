package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 18 {
		t.Errorf("Expected deck length of 18, but got %v", len(d))
		// t.Error("Expected deck length of 18, but got", len(d)) // will FAIL if use %v
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 18 {
		t.Errorf("Expected 18 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
