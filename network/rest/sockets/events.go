package sockets

type PlayerJoinEvent struct {
	Name   string      `json:"name"`
	Player interface{} `json:"players"`
}

type PlayerQuitEvent struct {
	Name   string      `json:"name"`
	Player interface{} `json:"player"`
}

type ServerCreateEvent struct {
	Name      string      `json:"name"`
	Container interface{} `json:"container"`
}
