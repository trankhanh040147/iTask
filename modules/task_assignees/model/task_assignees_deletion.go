package model

type TaskAssigneeDeletion struct {
	TaskId int `json:"task_id" form:"task_id"`
	UserId int `json:"user_id" form:"user_id"`
}

func (TaskAssigneeDeletion) TableName() string {
	return "TaskAssigned"
}

func (TaskAssigneeDeletion) GetEntityName() string {
	return EntityName
}
