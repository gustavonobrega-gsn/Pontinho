package game

import (
	"pontinho/internal/pontinho/config"
	"pontinho/internal/pontinho/entity"
)

type GameContext struct {
	gameConfig config.GameConfig
	deck       []entity.Card
}
