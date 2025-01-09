package main

func main() {

	//cards := newDeck()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
	//a, b := deal(cards, 3)
	//a.print()
	//b.print()

}
