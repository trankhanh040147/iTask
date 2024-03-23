package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task/model"
)

type CreateTaskStorage interface {
	CreateTask(ctx context.Context, data *model.TaskCreation) error
}

//
//type TaskTagStorage interface {
//	UpdateTaskTagsByTaskId(ctx context.Context, taskId int, tagIds []int) error
//	GetTaskTagsByTaskId(ctx context.Context, cond map[string]interface{}) (map[int]string, error)
//}
//
//type TagStorage interface {
//	CreateTagsByTagNames(ctx context.Context, tagType int, tags string) error
//	GetTagIdsByNames(ctx context.Context, tags string) ([]int, error)
//}

type createTaskBiz struct {
	store     CreateTaskStorage
	requester common.Requester
}

func NewCreateTaskBiz(store CreateTaskStorage, requester common.Requester) *createTaskBiz {
	return &createTaskBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createTaskBiz) CreateNewTask(ctx context.Context, data *model.TaskCreation) error {
	//if err := data.Validate(); err != nil {
	//	return common.ErrValidation(err)
	//}

	data.CreatedBy = biz.requester.GetUserId()

	if err := biz.store.CreateTask(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	//// create tags
	//if err := biz.tagStorage.CreateTagsByTagNames(ctx, int(tagModel.TypeTask), data.Tags); err != nil {
	//	return common.ErrCannotCreateEntity(tagModel.EntityName, err)
	//}
	//
	////// get tags id
	//TagIds, err := biz.tagStorage.GetTagIdsByNames(ctx, data.Tags)
	//if err != nil {
	//	return common.ErrCannotCreateEntity(tagModel.EntityName, err)
	//}
	//
	//// update task tags
	//if err := biz.taskTagStorage.UpdateTaskTagsByTaskId(ctx, data.Id, TagIds); err != nil {
	//	return common.ErrCannotCreateEntity(taskTagModel.EntityName, err)
	//}
	//
	//// !logging
	//fmt.Printf("<____________> Task w/ id %d has been created by user w/ id %d\n", data.Id, data.CreatedBy)
	//
	//// todo: add task_members
	//if err := biz.taskMemberStorage.CreateTaskMember(ctx, data.Id, data.CreatedBy); err != nil {
	//	return common.ErrCannotCreateEntity("task_members", err)
	//}
	//// todo: implement: worker(redis) | pubsub...

	return nil
}
