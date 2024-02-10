package iomodel

type ResetPassword struct {
	NewPassword string `json:"new_password" binding:"required"`
}
