package main

import "fmt"

func reversestring(card string) string {
	runecard := []rune(card)
	start := 0
	end := len(card) - 1
	for start < end {
		runecard[start], runecard[end] = runecard[end], runecard[start]
		start++
		end--
	}

	return string(runecard)
}

func main() {
	var card string
	fmt.Println("Enter card number: ")
	fmt.Scanln(&card)
	revcard := reversestring(card)
	fmt.Printf("%s\n", revcard)

}
