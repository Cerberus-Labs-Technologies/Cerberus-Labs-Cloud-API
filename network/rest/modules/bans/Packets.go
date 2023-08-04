package bans

type PlayerBannedPacket struct {
	Name string `json:"name"`
	Ban  Ban    `json:"ban"`
}

type PlayerUnbannedPacket struct {
	Name string `json:"name"`
	Ban  Ban    `json:"ban"`
}

type BanUpdatedPacket struct {
	Name string `json:"name"`
	Ban  Ban    `json:"ban"`
}
