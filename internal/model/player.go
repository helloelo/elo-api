package model

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

// Player is the model for player
type Player struct {
	IDPlayer     uint64 `db:"id_player" json:"id_player"`
	Name         string `db:"player" json:"player_name"`
	Picture      string `db:"picture" json:"picture"`
	Email        string `db:"email" json:"player_email"`
	Organization uint64 `db:"fk_organization" json:"id_organization"`
}

func (p Player) GetPlayerFromId(IDPlayer uint64) {
	qb := squirrel.Select("id_player").
		From("player").
		Where("id_player", IDPlayer)

	fmt.Println(qb.ToSql)
}
