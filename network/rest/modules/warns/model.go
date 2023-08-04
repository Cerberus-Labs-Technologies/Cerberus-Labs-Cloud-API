package warns

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type Warn struct {
	Id           int                `json:"id" db:"id"`
	WarnedPlayer string             `json:"warnedPlayer" db:"warned_player"`
	Executor     string             `json:"executor" db:"executor"`
	Reason       string             `json:"reason" db:"reason"`
	CreatedAt    database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt    database.TimeStamp `json:"updatedAt" db:"updated_at"`
}
