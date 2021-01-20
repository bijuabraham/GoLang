package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)
	//hand, remaining := deal(5, cards)
	//cards.print()
	//fmt.Println("My Hand:")
	//hand.print()
	//fmt.Println("\nDealer Hand:")
	//remaining.print()
	//fmt.Println(hand.toString())
	//fmt.Println(remaining.toString())
	//hand.saveDeck("hand.txt")
	//newHand := readDeck("hand.txt")
	//fmt.Println(newHand)
}
