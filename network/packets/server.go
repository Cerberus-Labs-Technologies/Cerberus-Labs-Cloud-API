package packets

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/network/io/groups"
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/network/io/templates"
)

type ServerCreatePacket struct {
	Name     string             `json:"name"`
	Template templates.Template `json:"template"`
	Group    groups.Group       `json:"group"`
}

type ServerActionPacket struct {
	Name     string `json:"name"`
	DockerId string `json:"dockerId"`
	Action   string `json:"action"` // start, stop, restart, delete
}
