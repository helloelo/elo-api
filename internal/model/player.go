package model

// Player is the model for player
type Player struct {
	IDPlayer     uint64 `db:"id_player" json:"id_player"`
	Name         string `db:"player" json:"player_name"`
	Email        string `db:"email" json:"player_email"`
	Organization uint64 `db:"fk_organization" json:"id_organization"`
}
