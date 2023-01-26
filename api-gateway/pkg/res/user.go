package res

type UserResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Error      string      `json:"error"`
	Data       interface{} `json:"user"`
	UserId     int64       `json:"user_id"`
	Token      string      `json:"token"`
}
