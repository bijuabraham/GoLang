package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuite+" of "+cardValue)
		}
	}
	return cards
}

func deal(handSize int, d deck) (deck, deck) {
	return d[:3], d[3:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeck(f string) {
	ioutil.WriteFile(f, []byte(d.toString()), 0777)
}
