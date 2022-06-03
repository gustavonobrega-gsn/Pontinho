package entity

import (
	gamestate "pontinho/internal/game-state"
)

type table struct {
	id      int64
	deck    []card
	players []player
	state   gamestate.GameState
}

var (
	tableId int64 = 0
)

func NewTable() *table {
	var tableDeck []card

	for _, cardLabel := range AllowedCardLabels {
		for _, cardSuit := range AllowedCardSuits {
			tableDeck = append(tableDeck, card{cardLabel, cardSuit})
		}
	}

	tableId++

	return &table{
		id:   tableId,
		deck: tableDeck}
}

func (t *table) GetDeck() []card {
	return t.deck
}
