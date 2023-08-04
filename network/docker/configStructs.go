package docker

type SpigotYml struct {
	Settings Settings `yaml:"settings"`
}

type Settings struct {
	Debug                     bool                 `yaml:"debug"`
	PlayerShuffle             int                  `yaml:"player-shuffle"`
	TimeoutTime               int                  `yaml:"timeout-time"`
	RestartOnCrash            bool                 `yaml:"restart-on-crash"`
	RestartScript             string               `yaml:"restart-script"`
	SaveUserCacheOnStopOnly   bool                 `yaml:"save-user-cache-on-stop-only"`
	SampleCount               int                  `yaml:"sample-count"`
	Bungeecord                bool                 `yaml:"bungeecord"`
	UserCacheSize             int                  `yaml:"user-cache-size"`
	MovedWronglyThreshold     float64              `yaml:"moved-wrongly-threshold"`
	MovedTooQuicklyMultiplier float64              `yaml:"moved-too-quickly-multiplier"`
	NettyThreads              int                  `yaml:"netty-threads"`
	Attribute                 map[string]Attribute `yaml:"attribute"`
	LogVillagerDeaths         bool                 `yaml:"log-villager-deaths"`
	LogNamedDeaths            bool                 `yaml:"log-named-deaths"`
}

type Attribute struct {
	Max float64 `yaml:"max"`
}

type WrapperConfig struct {
	APIKey         string `json:"apiKey"`
	CloudMasterUrl string `json:"cloudMasterUrl"`
	DockerId       string `json:"dockerId"`
}
