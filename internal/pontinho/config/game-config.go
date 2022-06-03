package config

var (
	defaultAllowedCardLabels = []string{"4", "5", "6", "7", "Q", "J", "K", "A", "1", "2", "3"}
	defaultAllowedCardSuits  = []string{"♦️", "♠️", "♥️", "♣️"}
)

type GameConfig struct {
	AllowedCardLabels []string
	AllowedCardSuits  []string
}

func NewGameConfig() *GameConfig {

	return &GameConfig{
		AllowedCardLabels: defaultAllowedCardLabels,
		AllowedCardSuits:  defaultAllowedCardSuits,
	}
}
