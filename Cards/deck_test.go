package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Deck length of 52 but got %v", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("First item in Deck should be Spades of Ace but got %v", d[0])
	}
	if d[len(d)-1] != "Hearts of Two" {
		t.Errorf("Last item in Deck should be Hearts of Two but got %v", d[len(d)-1])
	}
}
