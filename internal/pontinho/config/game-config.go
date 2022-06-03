package config

var (
	defaultAllowedCardLabels = []string{"4", "5", "6", "7", "Q", "J", "K", "A", "1", "2", "3"}
	defaultAllowedCardSuits  = []string{"♦️", "♠️", "♥️", "♣️"}
)

type GameConfig struct {
	allowedCardLabels []string
	allowedCardSuits  []string
}

func NewGameConfig() *GameConfig {

	return &GameConfig{
		allowedCardLabels: defaultAllowedCardLabels,
		allowedCardSuits:  defaultAllowedCardSuits,
	}
}
