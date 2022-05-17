package api

import (
	"fmt"
	"github.com/maldan/gam-app-audio_player/internal/app/audio_player/core"
	"github.com/maldan/go-cmhp/cmhp_convert"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-rapi/rapi_core"
	"strings"
)

type MainApi struct {
}

func (r MainApi) GetIndex(args ArgsEmpty) string {
	return "Test"
}

func (r MainApi) GetList() []core.AudioTrack {
	out := make([]core.AudioTrack, 0)

	list, _ := cmhp_file.ListAll(core.AppConfig.TrackDir)
	for _, file := range list {
		if !(strings.Contains(file.Name, ".mp3") || strings.Contains(file.Name, ".mp4")) {
			continue
		}

		out = append(out, core.AudioTrack{
			Name: file.Name,
			Path: file.FullPath,
			Url:  fmt.Sprintf("//%v:%v/api/main/track?path=%v", core.Host, core.Port, cmhp_convert.ToBase64(file.FullPath)),
		})
	}

	return out
}

func (r MainApi) GetTrack(ctx *rapi_core.Context, args ArgsTrack) string {
	ctx.IsServeFile = true
	return string(cmhp_convert.FromBase64(args.Path))
}
