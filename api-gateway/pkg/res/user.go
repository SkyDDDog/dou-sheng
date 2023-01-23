package res

type UserResponse struct {
	StatusCode int32       `json:"status_code"`
	Data       interface{} `json:"user"`
	StatusMsg  string      `json:"status_msg"`
	Error      string      `json:"error"`
	UserId     int64       `json:"user_id"`
	Token      string      `json:"token"`
}
