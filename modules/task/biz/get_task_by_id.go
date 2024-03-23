package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task/model"
)

type GetTaskStorage interface {
	GetTask(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Task, error)
}

type getTaskBiz struct {
	store GetTaskStorage
}

func NewGetTaskBiz(store GetTaskStorage) *getTaskBiz {
	return &getTaskBiz{store: store}
}

func (biz *getTaskBiz) GetTaskById(ctx context.Context, id int) (*model.Task, error) {
	data, err := biz.store.GetTask(ctx, map[string]interface{}{"id": id}, "Owner")

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	//// get tags

	//cond := map[string]interface{}{"id": id}

	//TaskTagsMap, err := biz.TaskTagStorage.GetTaskTagsByTaskId(ctx, cond)
	//if err != nil {
	//	return data, nil
	//}
	//
	//data.Tags = TaskTagsMap[data.Id]

	data.Parsing()

	return data, nil
}
