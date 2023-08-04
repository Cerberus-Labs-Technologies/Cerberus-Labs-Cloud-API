package player

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type IPlayer interface {
	GetUUID() UUID
	GetUsername() string
	GetXBoxUUID() database.JsonString
	GetLastIP() string
	GetLastSeen() database.TimeStamp
	GetFirstSeen() database.TimeStamp
}
