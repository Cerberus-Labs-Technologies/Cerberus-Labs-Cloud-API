package settings

type SettingCreatedPacket struct {
	Name    string  `json:"name"`
	Setting Setting `json:"setting"`
}

type SettingUpdatedPacket struct {
	Name       string  `json:"name"`
	OldSetting Setting `json:"oldSetting"`
	NewSetting Setting `json:"newSetting"`
}

type SettingDeletedPacket struct {
	Name    string  `json:"name"`
	Setting Setting `json:"setting"`
}
