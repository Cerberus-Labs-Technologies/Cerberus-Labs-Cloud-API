package permissions

type PermissionCreatedPacket struct {
	Name       string     `json:"name"`
	Permission Permission `json:"permission"`
}

type PermissionDeletedPacket struct {
	Name       string     `json:"name"`
	Permission Permission `json:"permission"`
}

type PermissionGroupCreatedPacket struct {
	Name            string          `json:"name"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
}

type PermissionGroupDeletedPacket struct {
	Name            string          `json:"name"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
}

type PermissionGroupUpdatedPacket struct {
	Name               string          `json:"name"`
	OldPermissionGroup PermissionGroup `json:"oldPermissionGroup"`
	NewPermissionGroup PermissionGroup `json:"newPermissionGroup"`
}

type PermissionRemovedFromGroupPacket struct {
	Name            string          `json:"name"`
	Permission      Permission      `json:"permission"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
}

type PermissionAddedToGroupPacket struct {
	Name            string          `json:"name"`
	Permission      Permission      `json:"permission"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
}

type PermissionAddedToPlayerPacket struct {
	Name             string           `json:"name"`
	PermissionPlayer PermissionPlayer `json:"permissionPlayer"`
	Permission       Permission       `json:"permission"`
}

type PermissionRemovedFromPlayerPacket struct {
	Name             string           `json:"name"`
	PermissionPlayer PermissionPlayer `json:"permissionPlayer"`
	Permission       Permission       `json:"permission"`
}

type PermissionGroupSetForPlayerPacket struct {
	Name            string          `json:"name"`
	PermissionGroup PermissionGroup `json:"permissionGroup"`
	PlayerUUID      string          `json:"playerUUID"`
}
