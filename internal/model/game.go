package model

// Game is the model for game
type Game struct {
	IDGame         uint64 `db:"id_game" json:"id_game"`
	Name           string `db:"name" json:"game_name"`
	Picture        string `db:"picture" json:"picture"`
	IDOrganization uint64 `db:"fk_organization" json:"id_organization"`
}
