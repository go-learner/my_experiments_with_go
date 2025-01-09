package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected length of 9 instead got %v", len(d))
	}
	if d[0] != "ace of club" {
		t.Errorf("Expected ace of club instead got %v", d[0])
	}
}

func TestNewDeck2(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadeddeck := newDeckFromFile("_decktesting")

	if len(loadeddeck) != 9 {
		t.Errorf("Expected length of 9 instead got %v", len(loadeddeck))
	}
	os.Remove("_decktesting")
}
