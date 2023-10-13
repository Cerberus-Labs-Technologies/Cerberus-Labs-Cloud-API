package inventory

type InventorySyncedPacket struct {
	Name          string        `json:"name"`
	UUID          string        `json:"uuid"`
	InventorySync InventorySync `json:"inventorySync"`
}
