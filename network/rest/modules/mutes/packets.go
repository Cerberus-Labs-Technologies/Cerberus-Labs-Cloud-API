package mutes

type PlayerMutedPacket struct {
	Name string `json:"name"`
	Mute Mute   `json:"mutes"`
}

type PlayerUnmutedPacket struct {
	Name string `json:"name"`
	Mute Mute   `json:"mutes"`
}

type MuteUpdatedPacket struct {
	Name string `json:"name"`
	Mute Mute   `json:"mutes"`
}
