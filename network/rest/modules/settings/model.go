package settings

type Setting struct {
	// ID is the primary key of the settings table.
	ID int `json:"id" db:"id"`
	// Name is the name of the setting.
	Name string `json:"settings_name" db:"settings_name"`
	// Value is the value of the setting.
	Value string `json:"value" db:"value"`
}
