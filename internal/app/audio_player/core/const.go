package core

type AudioTrack struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

type Config struct {
	TrackDir string `json:"trackDir"`
}

var Host = ""
var Port = 0
var DataDir = ""
var AppConfig Config
