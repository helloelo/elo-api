package model

// Elo is the model for elo
type Elo struct {
	IDPlayer uint64 `db:"fk_player" json:"id_player"`
	IDGame   uint64 `db:"fk_game" json:"id_game"`
	Value    uint64 `db:"value" json:"elo_value"`
}
