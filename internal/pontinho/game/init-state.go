package game

import (
	"pontinho/internal/pontinho/entity"
)

// InitState initial game state machine
type InitState struct {
	gameCtx *Context
}

// NewInitState creates a new init state
func NewInitState(gameContx *Context) State {

	state := &InitState{
		gameCtx: gameContx,
	}

	gameContx.state = state

	for _, cardLabel := range gameContx.gameConfig.AllowedCardLabels {
		for _, cardSuit := range gameContx.gameConfig.AllowedCardSuits {
			state.gameCtx.deck = append(state.gameCtx.deck, entity.Card{Label: cardLabel, Suit: cardSuit})
		}
	}

	return state
}
