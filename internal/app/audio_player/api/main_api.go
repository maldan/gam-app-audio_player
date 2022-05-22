package api

import (
	"fmt"
	"github.com/maldan/gam-app-audio_player/internal/app/audio_player/core"
	"github.com/maldan/go-cmhp/cmhp_convert"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
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

		d, _ := cmhp_process.Exec("ffprobe", "-v", "quiet", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", file.FullPath)
		d = strings.ReplaceAll(strings.Split(d, "\n")[0], "\r", "")
		duration := cmhp_convert.StrToFloat(d)

		out = append(out, core.AudioTrack{
			Name:     file.Name,
			Path:     file.FullPath,
			Url:      fmt.Sprintf("//%v:%v/api/main/track?path=%v", core.Host, core.Port, cmhp_convert.ToBase64(file.FullPath)),
			Size:     cmhp_file.Size(file.FullPath),
			Duration: duration,
		})
	}

	return out
}

func (r MainApi) GetTrack(ctx *rapi_core.Context, args ArgsTrack) string {
	ctx.IsServeFile = true
	return string(cmhp_convert.FromBase64(args.Path))
}

func (r MainApi) GetSettings() core.Settings {
	settings := core.Settings{}
	cmhp_file.ReadJSON(core.DataDir+"/settings.json", &settings)
	return settings
}

func (r MainApi) PostSettings(args core.Settings) {
	cmhp_file.Write(core.DataDir+"/settings.json", &args)
}
