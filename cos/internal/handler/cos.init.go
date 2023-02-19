package handler

import (
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var COS_CLIENT *cos.Client

func InitCos() {
	bucketUrl := GetBucketUrl()
	bu, _ := url.Parse(bucketUrl)
	serviceUrl := GetServiceUrl()
	su, _ := url.Parse(serviceUrl)
	b := &cos.BaseURL{
		BucketURL:  bu,
		ServiceURL: su,
	}
	COS_CLIENT = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("main.secretId"),
			SecretKey: viper.GetString("main.secretKey"),
		},
	})
}

func GetBucketUrl() string {
	bucket := viper.GetString("main.bucket")
	appid := viper.GetString("main.appid")
	region := viper.GetString("main.region")
	return "https://" + bucket + "-" + appid + ".main." + region + ".myqcloud.com"
}

func GetServiceUrl() string {
	region := viper.GetString("main.region")
	return "https://cos." + region + ".myqcloud.com"
}
