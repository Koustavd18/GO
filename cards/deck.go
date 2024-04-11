package main

import "fmt"

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuit := []string{"Spades", "Clubs", "Hearts", "Diamond"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
return cards
}

func ( d deck ) print() {
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}
