package reports

import "github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"

// Report represents a report in the database
type Report struct {
	ID              int                 `json:"id" db:"id"`
	ReportingPlayer database.JsonString `json:"reportingPlayer" db:"reporting_player"`
	ReportedPlayer  database.JsonString `json:"reportedPlayer" db:"reported_player"`
	Reason          database.JsonString `json:"reason" db:"reason"`
	Investigated    database.IntBool    `json:"investigated" db:"investigated"`
	InvestigatedBy  database.JsonString `json:"investigatedBy" db:"investigated_by"`
	CreatedAt       database.TimeStamp  `json:"createdAt" db:"created_at"`
	UpdatedAt       database.TimeStamp  `json:"updatedAt" db:"updated_at"`
	InvestigatedAt  database.TimeStamp  `json:"investigatedAt" db:"investigated_at"`
}
