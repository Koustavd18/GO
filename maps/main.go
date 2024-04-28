package main

import "fmt"

func main() {

	//var colors map[string] string

	// colore := make(map[string] string)

	// colore["white"] = "#ffffff"

	// colore["blue"] = "neel"

	colors := map[string] string {
		"red": "#ff000",
		"green" : "#4xfj0", 
		"white": "#ffffff",
	}

	// delete(colore, "white")

	printMap(colors)
	// fmt.Println(colore)
}

func printMap(c map[string] string) {
	for color, hex := range c {
		fmt.Println("Hex code for color ", color , "is ", hex)
	}
}