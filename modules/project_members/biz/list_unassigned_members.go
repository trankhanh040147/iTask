package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project_members/model"
	taskModel "iTask/modules/task/model"
	taskAssigneModel "iTask/modules/task_assignees/model"
)

type TaskStorage interface {
	GetTask(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*taskModel.Task, error)
}
type TaskAssigneeStorage interface {
	ListAssigneeByTaskId(ctx context.Context, taskId int, moreKeys ...string) ([]taskAssigneModel.TaskAssignee, error)
	GetAssignee(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*taskAssigneModel.TaskAssignee, error)
}

type listUnassignedMembersBiz struct {
	TaskAssigneeStore TaskAssigneeStorage
	TaskStore         TaskStorage
	store             ListMembersByIdStore
}

func NewListUnassignedMembersBiz(store ListMembersByIdStore, tas TaskAssigneeStorage, ts TaskStorage) *listUnassignedMembersBiz {
	return &listUnassignedMembersBiz{store: store, TaskAssigneeStore: tas, TaskStore: ts}
}

func (biz *listUnassignedMembersBiz) ListUnassignedMembers(ctx context.Context, projectId, taskId int) ([]model.SimpleMember, error) {
	// todo: check if projectId
	// todo: check is this tsk belong to this project

	// *check task is valid
	_, err := biz.TaskStore.GetTask(ctx, map[string]interface{}{"id": taskId})
	if err != nil {
		return nil, common.ErrCannotGetEntity("Task", err)
	}

	// 1. Get list user id of members that assigned to the task
	assignees, err := biz.TaskAssigneeStore.ListAssigneeByTaskId(ctx, taskId)
	if err != nil {
		return nil, common.ErrCannotListEntity("TaskAssignee", err)
	}
	mapAssignees := map[int]bool{}
	for _, assignee := range assignees {
		mapAssignees[assignee.UserId] = true
	}

	// 2. Get list member of projects
	members, err := biz.store.ListMembersById(ctx, projectId)
	if err != nil {
		return nil, common.ErrCannotListEntity("ProjectMember", err)
	}

	// 3. Filter out members that are not in the list of members assigned to the task
	var unassignedMembers []model.SimpleMember
	for _, member := range members {
		if !mapAssignees[member.UserId] {
			unassignedMembers = append(unassignedMembers, member)
		}
	}

	return unassignedMembers, nil
}
