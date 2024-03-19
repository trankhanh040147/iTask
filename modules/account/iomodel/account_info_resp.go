package iomodel

type AccountInfoResp struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Title    string `json:"title"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Bio      string `json:"bio"`
	Avt      string `json:"profile_ava_url"`
	Cover    string `json:"profile_cover_url"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
