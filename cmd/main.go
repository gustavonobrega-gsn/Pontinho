package main

import (
	"fmt"
	"pontinho/internal/entity"
)

func main() {

	table := entity.NewTable()

	deck := table.GetDeck()

	for _, element := range deck {
		fmt.Println(element.Label, element.Suit)
	}

}
