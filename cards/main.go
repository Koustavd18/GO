package main

func main() {
	// cards := newDeck()
	// hand, remainingCard := deal(cards, 5)
	// fmt.Println("Cards in hand")
	// hand.print()
	// fmt.Println("Cards Remaining")
	// remainingCard.print()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	}	
