package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, j := range d {
		fmt.Println(j)
	}
}

func newDeck() deck {
	cards := deck{}
	a := []string{"diamond", "heart", "clubs", "spades"}

	b := []string{"ace", "two", "three"}

	for _, i := range b {
		for _, k := range a {
			cards = append(cards, i+"of"+k)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {

	return d[:handsize], d[handsize:]

}
