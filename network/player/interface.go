package player

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type IPlayer interface {
	getUUID() UUID
	getUsername() string
	getXBoxUUID() database.JsonString
	getLastIP() string
	getLastSeen() database.TimeStamp
	getFirstSeen() database.TimeStamp
}
