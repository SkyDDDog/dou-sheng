package res

type MessageResponse struct {
	StatusCode int32       `json:"status_code"`
	Data       interface{} `json:"message_list"`
	StatusMsg  string      `json:"status_msg"`
	Error      string      `json:"error"`
}
