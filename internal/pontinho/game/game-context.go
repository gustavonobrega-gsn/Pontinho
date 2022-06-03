package game

import (
	"pontinho/internal/pontinho/config"
	"pontinho/internal/pontinho/entity"
)

// Context store game properties
type Context struct {
	gameConfig config.GameConfig
	deck       []entity.Card
	state      State
}
