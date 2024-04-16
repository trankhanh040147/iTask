package model

//var (
//	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
//)

const (
	EntityName = "ProjectMemberInvited"
)

type ProjectMemberInvited struct {
	Id                 int `json:"id" gorm:"column:id"`
	ProjectId          int `json:"project_id" gorm:"column:project_id"`
	VerificationMailId int `json:"mail_id" gorm:"column:email_verification_id"`
}

func (ProjectMemberInvited) TableName() string {
	return "ProjectMemberInvited"
}

func (a *ProjectMemberInvited) GetProjectID() int {
	return a.ProjectId
}
