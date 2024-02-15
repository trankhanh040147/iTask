package iomodel

type AccountRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}
