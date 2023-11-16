package cloudNpc

type CloudNpcCreatePacket struct {
	Name     string   `json:"name"`
	UUID     string   `json:"uuid"`
	CloudNPC CloudNPC `json:"cloudNPC"`
}

type CloudNpcRemovePacket struct {
	Name     string   `json:"name"`
	UUID     string   `json:"uuid"`
	CloudNPC CloudNPC `json:"cloudNPC"`
}

type CloudNpcUpdatePacket struct {
	Name     string   `json:"name"`
	UUID     string   `json:"uuid"`
	CloudNPC CloudNPC `json:"cloudNPC"`
}
