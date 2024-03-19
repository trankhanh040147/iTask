package iomodel

type AccountInfoToAdminManageResp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Avt      string `json:"profile_ava_url"`
	Cover    string `json:"profile_cover_url"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
