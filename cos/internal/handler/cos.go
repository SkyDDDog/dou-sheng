package handler

import (
	"bytes"
	"context"
	"cos/internal/service"
	"cos/pkg/e"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

type CosService struct {
	service.UnimplementedCosServiceServer
}

func NewCosService() *CosService {
	return &CosService{}
}

func (*CosService) UploadFile(ctx context.Context, req *service.CosUploadRequest) (resp *service.CosUploadResponse, err error) {
	resp = new(service.CosUploadResponse)
	video, err := UploadVideo(req.Id, req.Data)
	resp.VideoUrl = video
	snapshot, err := GetSnapshot(req.GetId(), 1)
	fmt.Println("snapshot: ", snapshot)
	photo, err := UploadPhoto(req.Id, snapshot)
	resp.CoverUrl = photo
	if err != nil {
		resp.StatusCode = e.ErrorUploadFail
		resp.StatusMsg = e.GetMsg(e.ErrorUploadFail)
	}

	return resp, err
}

func UploadPhoto(id int64, data []byte) (url string, err error) {
	file := bytes.NewReader(data)

	key := fmt.Sprintf("cover/%d.jpg", id)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
	}
	_, err = COS_CLIENT.Object.Put(context.Background(), key, file, opt)
	photoUrl := COS_CLIENT.Object.GetObjectURL(key)
	return GetBucketUrl() + photoUrl.Path, err
}

func UploadVideo(id int64, data []byte) (url string, err error) {
	file := bytes.NewReader(data)

	key := fmt.Sprintf("video/%d.mp4", id)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "video/mp4",
		},
	}
	_, err = COS_CLIENT.Object.Put(context.Background(), key, file, opt)
	photoUrl := COS_CLIENT.Object.GetObjectURL(key)
	return GetBucketUrl() + photoUrl.Path, err
}

func GetSnapshot(id int64, frameNum int) (data []byte, err error) {
	buf := bytes.NewBuffer(nil)
	videoUrl := GetBucketUrl() + fmt.Sprintf("/video/%d.mp4", id)
	err = ffmpeg.Input(videoUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	return buf.Bytes(), err
}
