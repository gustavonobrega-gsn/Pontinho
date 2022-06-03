package game

import (
	"pontinho/internal/pontinho/entity"
)

type initState struct {
	gameCtx *GameContext
}

func NewInitState(gameContx *GameContext) *initState {

	return &initState{
		gameCtx: gameContx,
	}

}

func (state *initState) init() {

	for _, cardLabel := range entity.AllowedCardLabels {
		for _, cardSuit := range entity.AllowedCardSuits {
			state.gameCtx.deck = append(state.gameCtx.deck, entity.Card{Label: cardLabel, Suit: cardSuit})
		}
	}

}
