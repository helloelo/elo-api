package model

import "time"

// Match is the model for matches
type Match struct {
	IDMatch      uint64    `db:"id_match" json:"id_match"`
	IDGame       uint64    `db:"fk_game" json:"id_game"`
	IDStatus     uint64    `db:"id_status" json:"id_status"`
	Player1      []Player  `db:"player1" json:"player1"`
	Player2      []Player  `db:"player2" json:"player2"`
	Player1Score int       `db:"player1_score" json:"player1_score"`
	Player2Score int       `db:"player2_score" json:"player2_score"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

const (
	AwaitingConfirmation = 1
	Declined             = 2
	Validated            = 3
	AwaitingApproval     = 4
	Completed            = 5
	Invalidated          = 6
)
