package worker

import (
	"iTask/constant"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
	UrlVerifyEmail      = constant.URL_HOST_EC2 + "/verify_email"
)

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}
