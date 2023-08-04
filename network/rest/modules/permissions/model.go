package permissions

import (
	"github.com/Cerberus-Labs-Technologies/Cerberus-Labs-Cloud-API/database"
)

type Permission struct {
	Id             int                `json:"id" db:"id"`
	PermissionName string             `json:"permissionName" db:"permission_name"`
	CreatedAt      database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdatedAt      database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

type PermissionGroup struct {
	Id         int                 `json:"id" db:"id"`
	GroupName  string              `json:"groupName" db:"group_name"`
	Prefix     database.JsonString `json:"prefix" db:"prefix"`
	Suffix     database.JsonString `json:"suffix" db:"suffix"`
	Color      database.JsonString `json:"color" db:"color"`
	Priority   int                 `json:"priority" db:"priority"`
	Fallback   database.IntBool    `json:"fallback" db:"fallback"`
	CreatedAt  database.TimeStamp  `json:"createdAt" db:"created_at"`
	UpdatedAt  database.TimeStamp  `json:"updatedAt" db:"updated_at"`
	Permission []Permission        `json:"permissions"` // This is a virtual field
}

type PermissionGroupPermission struct {
	Id                int                `json:"id" db:"id"`
	PermissionGroupId int                `json:"permissionGroupId" db:"permission_group_id"`
	PermissionId      int                `json:"permissionId" db:"permission_id"`
	CreatedAt         database.TimeStamp `json:"createdAt" db:"created_at"`
}

type PermissionGroupUser struct {
	Id                int                `json:"id" db:"id"`
	PermissionGroupId int                `json:"permissionGroupId" db:"permission_group_id"`
	UserUUID          string             `json:"userUUID" db:"user_uuid"`
	CreatedAt         database.TimeStamp `json:"createdAt" db:"created_at"`
	UpdateAt          database.TimeStamp `json:"updatedAt" db:"updated_at"`
}

type PermissionPlayer struct {
	Id           int    `json:"id" db:"id"`
	PermissionId int    `json:"permissionId" db:"permission_id"`
	UserUUID     string `json:"userUUID" db:"user_uuid"`
}

type PermPlayer struct {
	Uuid            string          `json:"uuid"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
	Permissions     []Permission    `json:"permissions"` // This is a virtual field
}
