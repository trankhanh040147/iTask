package model

import (
	"errors"
	"iTask/common"
	userModel "iTask/modules/account/iomodel"
	"iTask/modules/project/model"
	taskAssigneesModel "iTask/modules/task_assignees/model"
	"time"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrTaskIsDeleted     = errors.New("task is deleted")
)

const (
	EntityName = "Task"
)

type Task struct {
	common.SQLModel
	Name          string                               `json:"name" gorm:"column:name"`
	Description   string                               `json:"description" gorm:"column:description"`
	Status        TaskStatus                           `json:"status" gorm:"column:status"`
	StatusValue   string                               `json:"-" gorm:"-"`
	ProjectId     int                                  `json:"-" gorm:"column:project_id"`
	CreatedBy     int                                  `json:"-" gorm:"column:created_by"`
	ParentTask    int                                  `json:"parent_task_id" gorm:"column:parent_task_id"`
	Position      float64                              `json:"position" gorm:"column:position"`
	Priority      TaskPriority                         `json:"priority" gorm:"column:priority"`
	PriorityValue string                               `json:"-" gorm:"-"`
	Completed     bool                                 `json:"-" gorm:"column:completed"`
	DueDate       *time.Time                           `json:"due_date" gorm:"column:due_date"`
	StartedAt     *time.Time                           `json:"started_at" gorm:"column:started_at"`
	CompletedAt   *time.Time                           `json:"completed_at" gorm:"column:completed_at"`
	ProjectInfo   *model.SimpleProject                 `json:"project_info" gorm:"foreignKey:ProjectId"`
	Owner         *userModel.SimpleAccount             `json:"owner" gorm:"foreignKey:CreatedBy"`
	Assignees     *[]taskAssigneesModel.SimpleAssignee `json:"assignees" gorm:"foreignKey:TaskId"`
}

func (Task) TableName() string {
	return "Tasks"
}

func (t *Task) GetStatus() string {
	return t.Status.String()
}

func (t *Task) GetCreatedBy() int {
	return t.CreatedBy
}

func (t *Task) GetID() int {
	return t.Id
}

func (t *Task) Parsing() {
	t.StatusValue = t.Status.String()
	t.PriorityValue = t.Priority.String()
}

type TaskStatus int
type TaskPriority int

const (
	StatusPending TaskStatus = 1 + iota
	StatusInProgress
	StatusDone
	StatusDeleted
)
const (
	PriorityHigh TaskPriority = 1 + iota
	PriorityMedium
	PriorityLow
)

func (status TaskStatus) String() string {
	switch status {
	case StatusPending:
		return "Pending"
	case StatusInProgress:
		return "In Progress"
	case StatusDone:
		return "Done"
	case StatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func (priority TaskPriority) String() string {
	switch priority {
	case PriorityHigh:
		return "High"
	case PriorityMedium:
		return "Medium"
	case PriorityLow:
		return "Low"
	default:
		return "Unknown"
	}
}
