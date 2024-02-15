package iomodel

type AccountInfoToAdminManageResp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Avt      string `json:"profile_ava_url"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
