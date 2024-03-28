package model

type ProjectMemberInvitation struct {
	ProjectId int    `json:"project_id" form:"project_id"`
	UserEmail string `json:"email" form:"email"`
}
