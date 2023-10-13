package inventory

import "github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"

type InventorySync struct {
	ID        int                `json:"id" db:"id"`
	UUID      string             `json:"uuid" db:"uuid"`
	Inventory string             `json:"inventory" db:"inventory"`
	CreatedAt database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt database.TimeStamp `json:"updatedAt" db:"updated_at"`
}
