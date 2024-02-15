package iomodel

type AccountUpdatePersonalInfo struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Avt      string `json:"avatar"`
	Bio      string `json:"bio"`
}
