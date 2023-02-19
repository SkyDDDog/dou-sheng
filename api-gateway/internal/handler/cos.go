package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func UploadPhoto(ginCtx *gin.Context) {
	var req service.CosUploadRequest
	cosService := ginCtx.Keys["main"].(service.CosServiceClient)
	data, err := ginCtx.FormFile("data")
	id, _ := strconv.ParseInt(ginCtx.PostForm("id"), 10, 64)
	req.Id = id
	buff := new(bytes.Buffer)
	PanicIfCosError(err)
	dataContent, _ := data.Open()
	io.Copy(buff, dataContent)
	req.Data = buff.Bytes()
	PanicIfCosError(err)
	resp, err := cosService.UploadFile(context.Background(), &req)
	PanicIfCosError(err)
	fmt.Println("cover: ", resp.CoverUrl)
	fmt.Println("video: ", resp.VideoUrl)
	r := res.Response{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		Error:      "",
	}
	ginCtx.JSON(http.StatusOK, r)
}
