package main

import "fmt"

func main() {
	var variable = []string{"hello", "worlds", "good", "morning" }
	fmt.Println(variable[2:])
	fmt.Println(variable[:2])
}
