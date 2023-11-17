package cloudNpc

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/network/io/groups"
	"github.com/google/uuid"
)

type CloudNPC struct {
	Name               string               `json:"name" db:"name"`
	UUID               uuid.UUID            `json:"uuid" db:"uuid"`
	Skin               string               `json:"skin" db:"skin"`
	Group              groups.Group         `json:"group" db:"group"`
	Location           Location             `json:"location" db:"location"`
	Sneaking           database.IntBool     `json:"sneaking" db:"sneaking"`
	CopiesPlayer       database.IntBool     `json:"copiesPlayer" db:"copies_player"`
	LookAtPlayer       database.IntBool     `json:"lookAtPlayer" db:"look_at_player"`
	Visible            database.IntBool     `json:"visible" db:"visible"`
	Glowing            database.IntBool     `json:"glowing" db:"glowing"`
	BouncePlayers      database.IntBool     `json:"bouncePlayers" db:"bounce_players"`
	BouncePlayersRange database.JsonFloat64 `json:"bouncePlayersRange" db:"bounce_players_range"`
}

type Location struct {
	World string               `json:"world" db:"world"`
	X     database.JsonFloat64 `json:"x" db:"x"`
	Y     database.JsonFloat64 `json:"y" db:"y"`
	Z     database.JsonFloat64 `json:"z" db:"z"`
	Yaw   database.JsonFloat64 `json:"yaw" db:"yaw"`
	Pitch database.JsonFloat64 `json:"pitch" db:"pitch"`
}
