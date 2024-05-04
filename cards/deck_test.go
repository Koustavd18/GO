package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d:= newDeck()

	if (len(d) != 24){
		t.Errorf("Supposed to recieve 24. But recieved %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile (t *testing.T){
	os.Remove("_deckTesting")

	deck:= newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck:= newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 24 {
		t.Errorf("Error loading from file. Recieved length %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}