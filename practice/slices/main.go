package main

func main() {

	cards := newDeck()
	hand, remainingDeck := deal(cards, 2)
	hand.print()
	remainingDeck.print()

}
