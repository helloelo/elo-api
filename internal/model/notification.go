package model

import "database/sql"

// Notification is the model related to notifications
type Notification struct {
	IDNotification     uint64       `db:"id_notification" json:"id_notification"`
	IDNotificationType int          `db:"fk_notification_type" json:"id_notification_type"`
	Accepted           sql.NullBool `db:"accepted" json:"accepted"`
}

const (
	// NotificationChallenge is the id of notification related to challenge accept / decline
	NotificationChallenge = 1
	// NotificationApproval is the id of notification related to score accept / decline
	NotificationApproval = 2
)
