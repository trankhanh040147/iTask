package iomodel

type AccountUpdatePersonalInfo struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Title    string `json:"title"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
	Avt      string `json:"profile_ava_url"`
	Cover    string `json:"profile_cover_url"`
	Bio      string `json:"bio"`
}
