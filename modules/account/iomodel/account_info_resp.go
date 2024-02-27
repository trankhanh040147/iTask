package iomodel

type AccountInfoResp struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Bio      string `json:"bio"`
	Avt      string `json:"profile_ava_url"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
