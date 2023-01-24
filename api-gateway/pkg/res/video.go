package res

type VideoResponse struct {
	StatusCode int32       `json:"status_code"`
	Data       interface{} `json:"video_list"`
	StatusMsg  string      `json:"status_msg"`
	Error      string      `json:"error"`
}

type FeedResponse struct {
	StatusCode int32       `json:"status_code"`
	Data       interface{} `json:"video_list"`
	NextTime   int64       `json:"next_time"`
	StatusMsg  string      `json:"status_msg"`
	Error      string      `json:"error"`
}
