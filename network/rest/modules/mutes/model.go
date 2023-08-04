package mutes

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type Mute struct {
	ID          int                `json:"id" db:"id"`
	MutedPlayer string             `json:"mutedPlayer" db:"muted_player"`
	Executor    string             `json:"executor" db:"executor"`
	Permanent   database.IntBool   `json:"permanent" db:"permanent"`
	Until       database.TimeStamp `json:"until" db:"until"`
	CreatedAt   database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt   database.TimeStamp `json:"updatedAt" db:"updated_at"`
}
