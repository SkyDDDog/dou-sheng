package res

type FavoriteResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"msg"`
	Error      string      `json:"Error"`
	Data       interface{} `json:"video_list"`
}

type CommentResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"msg"`
	Error      string      `json:"Error"`
	Data       interface{} `json:"comment"`
}

type CommentListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"msg"`
	Error      string      `json:"Error"`
	Data       interface{} `json:"comment_list"`
}
