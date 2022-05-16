package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-audio_player/internal/app/audio_player"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
