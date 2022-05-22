package core

type AudioTrack struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Url      string  `json:"url"`
	Duration float64 `json:"duration"`
	Size     int64   `json:"size"`
}

type Config struct {
	TrackDir string `json:"trackDir"`
}

type Settings struct {
	TrackId     int     `json:"trackId"`
	Volume      float64 `json:"volume"`
	CurrentTime float64 `json:"currentTime"`
}

var Host = ""
var Port = 0
var DataDir = ""
var AppConfig Config
