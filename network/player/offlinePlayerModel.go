package player

import (
	"encoding/json"
	database "github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type OfflinePlayer struct {
	ID        int                 `json:"id" db:"id"`
	UUID      UUID                `json:"uuid" db:"uuid"`
	Username  string              `json:"username" db:"username"`
	XBoxUUID  database.JsonString `json:"xboxUUID" db:"xbox_uuid"`
	LastIP    string              `json:"lastIP" db:"last_ip"`
	LastSeen  database.TimeStamp  `json:"lastSeen" db:"last_seen"`
	FirstSeen database.TimeStamp  `json:"firstSeen" db:"first_seen"`
	// Rank        string `json:"rank"` // TODO: add rank
}

func (p *OfflinePlayer) getUUID() UUID {
	return p.UUID
}

func (p *OfflinePlayer) getUsername() string {
	return p.Username
}

func (p *OfflinePlayer) getXBoxUUID() string {
	return p.XBoxUUID.String
}

func (p *OfflinePlayer) getLastIP() string {
	return p.LastIP
}

func (p *OfflinePlayer) getLastSeen() database.TimeStamp {
	return p.LastSeen
}

func (p *OfflinePlayer) getFirstSeen() database.TimeStamp {
	return p.FirstSeen
}

func (p *OfflinePlayer) FromJSON(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		return err
	}
	return nil
}
