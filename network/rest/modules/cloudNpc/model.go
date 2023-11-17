package cloudNpc

import (
	"encoding/json"
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

func (cp *CloudNPC) ToCloudNPCDB() CloudNPCDB {
	return CloudNPCDB{
		Name:               cp.Name,
		UUID:               cp.UUID,
		Skin:               cp.Skin,
		GroupName:          cp.Group.Name,
		Location:           cp.Location.ToJsonString(),
		Sneaking:           false,
		CopiesPlayer:       false,
		LookAtPlayer:       false,
		Visible:            false,
		Glowing:            false,
		BouncePlayers:      false,
		BouncePlayersRange: cp.BouncePlayersRange,
	}
}

type CloudNPCDB struct {
	Name               string               `json:"name" db:"name"`
	UUID               uuid.UUID            `json:"uuid" db:"uuid"`
	Skin               string               `json:"skin" db:"skin"`
	GroupName          string               `json:"groupName" db:"group_name"`
	Location           string               `json:"location" db:"location"`
	Sneaking           database.IntBool     `json:"sneaking" db:"sneaking"`
	CopiesPlayer       database.IntBool     `json:"copiesPlayer" db:"copies_player"`
	LookAtPlayer       database.IntBool     `json:"lookAtPlayer" db:"look_at_player"`
	Visible            database.IntBool     `json:"visible" db:"visible"`
	Glowing            database.IntBool     `json:"glowing" db:"glowing"`
	BouncePlayers      database.IntBool     `json:"bouncePlayers" db:"bounce_players"`
	BouncePlayersRange database.JsonFloat64 `json:"bouncePlayersRange" db:"bounce_players_range"`
}

func (cp *CloudNPCDB) ToCloudNPC(groupss []groups.Group) (CloudNPC, error) {
	location := Location{}
	err := location.FromJsonString(cp.Location)
	if err != nil {
		return CloudNPC{}, err
	}
	var group groups.Group
	for _, g := range groupss {
		if g.Name == cp.GroupName {
			group = g
			break
		}
	}
	if group.Name == "" {
		return CloudNPC{}, nil
	}
	return CloudNPC{
		Name:               cp.Name,
		UUID:               cp.UUID,
		Skin:               cp.Skin,
		Group:              group,
		Location:           location,
		Sneaking:           cp.Sneaking,
		CopiesPlayer:       cp.CopiesPlayer,
		LookAtPlayer:       cp.LookAtPlayer,
		Visible:            cp.Visible,
		Glowing:            cp.Glowing,
		BouncePlayers:      cp.BouncePlayers,
		BouncePlayersRange: cp.BouncePlayersRange,
	}, nil
}

type Location struct {
	World string               `json:"world"`
	X     database.JsonFloat64 `json:"x"`
	Y     database.JsonFloat64 `json:"y"`
	Z     database.JsonFloat64 `json:"z"`
	Yaw   database.JsonFloat64 `json:"yaw"`
	Pitch database.JsonFloat64 `json:"pitch"`
}

func (l *Location) ToJsonString() string {
	data, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(data)
}

func (l *Location) FromJsonString(data string) error {
	return json.Unmarshal([]byte(data), &l)
}
