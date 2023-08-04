package warns

type PlayerWarnedPacket struct {
	Name string `json:"name"`
	Warn Warn   `json:"warn"`
}

type PlayerUnwarnedPacket struct {
	Name string `json:"name"`
	Warn Warn   `json:"warn"`
}

type WarnUpdatedPacket struct {
	Name string `json:"name"`
	Warn Warn   `json:"warn"`
}
