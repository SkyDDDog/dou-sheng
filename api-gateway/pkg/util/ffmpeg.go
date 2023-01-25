package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"strconv"
)

func GetSnapshot(videoId int64, frameNum int) error {
	videoSuffix := ".mp4"
	snapshotSuffix := ".png"
	snapshotPath := "./public/cover/" + strconv.FormatInt(videoId, 10) + snapshotSuffix
	videoPath := "./public/video/" + strconv.FormatInt(videoId, 10) + videoSuffix
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	err = imaging.Save(img, snapshotPath)
	if err != nil {
		return err
	}
	return nil
}
