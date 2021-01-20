package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remaining := deal(5, cards)
	//cards.print()
	//fmt.Println("My Hand:")
	//hand.print()
	//fmt.Println("\nDealer Hand:")
	//remaining.print()
	fmt.Println(hand.toString())
	fmt.Println(remaining.toString())
	hand.saveDeck("hand.txt")
}
