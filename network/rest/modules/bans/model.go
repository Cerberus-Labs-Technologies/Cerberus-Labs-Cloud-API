package bans

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type Ban struct {
	Id           int                `json:"id" db:"id"`
	BannedPlayer string             `json:"bannedPlayer" db:"banned_player"`
	BannedBy     string             `json:"bannedBy" db:"banned_by"`
	Reason       string             `json:"reason" db:"reason"`
	Permanent    database.IntBool   `json:"permanent" db:"permanent"`
	BanHash      string             `json:"banHash" db:"ban_hash"`
	ExpiresAt    database.TimeStamp `json:"expiresAt" db:"expires_at"`
	CreatedAt    database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt    database.TimeStamp `json:"updatedAt" db:"updated_at"`
	Active       database.IntBool   `json:"active" db:"active"`
}
