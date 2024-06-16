package ffmpeghlsstreamer

import (
	"context"
	"fmt"
	utils "hls-streamer/shell_utils"
	"time"

	"go.uber.org/zap"
)

const FFMPEG_BINARY_PATH = "/usr/bin/ffmpeg"

type HlsStreamerStruct struct {
	logger       *zap.SugaredLogger
	InputUrlChan chan string
}

func (hlsStreamer *HlsStreamerStruct) Init(logger *zap.SugaredLogger, hlsInputUrlChan chan string) {
	hlsStreamer.logger = logger
	hlsStreamer.InputUrlChan = hlsInputUrlChan
	hlsStreamer.logger.Info("Successfully initialised hlsStreamer")
	go hlsStreamer.handle()

}

// ffmpeg  -i udp://127.0.0.1:2222 -c copy -f hls -hls_time 5 playlist.m3u8
func (HlsStreamer *HlsStreamerStruct) GetShellCommandArgs(hlsInputUrl string) string {
	shellCommandargs := fmt.Sprintf("-i %s -c copy -f hls -hls_time 4 /home/rohanb/rohan_projects/go_projects/hls-streaming-app/hls-stream-files/playlist.m3u8", hlsInputUrl)
	return shellCommandargs

}
func (hlsStreamer *HlsStreamerStruct) handle() {
	var cancelFunc context.CancelFunc
	ctx := context.Background()
	for {
		select {
		case inputUrl := <-hlsStreamer.InputUrlChan:
			hlsStreamer.logger.Infof("Received new hlsInputUrl %s", inputUrl)
			if cancelFunc != nil {
				cancelFunc()
			}
			ctx, cancelFunc = context.WithCancel(ctx)
			hlsStreamerInfoStruct := utils.HlsStreamInfoStruct{}
			hlsStreamerInfoStruct.Command = FFMPEG_BINARY_PATH
			hlsStreamerInfoStruct.CmdExecPath = hlsStreamer.GetShellCommandArgs(inputUrl)
			hlsStreamer.logger.Infof("Running the hls streamer command %s in %s", hlsStreamerInfoStruct.Command, hlsStreamerInfoStruct.CmdExecPath)
			go utils.ExecuteCommandWithUpdates(hlsStreamerInfoStruct)
		default:
			time.Sleep(10 * time.Second)

		}
	}
}
