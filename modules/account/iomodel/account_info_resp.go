package iomodel

type AccountInfoResp struct {
	Id       int    `json:"id"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Bio      string `json:"bio"`
	Avt      string `json:"avatar"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
