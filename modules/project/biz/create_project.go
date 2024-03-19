package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
	projectTagModel "iTask/modules/project_tags/model"
	tagModel "iTask/modules/tag/model"
)

type CreateProjectStorage interface {
	CreateProject(ctx context.Context, data *model.ProjectCreation) error
}

type ProjectTagStorage interface {
	UpdateProjectTagsByProjectId(ctx context.Context, projectId int, tagIds []int) error
	GetProjectTagsByProjectId(ctx context.Context, cond map[string]interface{}) (map[int]string, error)
}

type TagStorage interface {
	CreateTagsByTagNames(ctx context.Context, tagType int, tags string) error
	GetTagIdsByNames(ctx context.Context, tags string) ([]int, error)
}

type createProjectBiz struct {
	store             CreateProjectStorage
	projectTagStorage ProjectTagStorage
	tagStorage        TagStorage
}

func NewCreateProjectBiz(store CreateProjectStorage, projectTagStorage ProjectTagStorage, tagStorage TagStorage) *createProjectBiz {
	return &createProjectBiz{store: store, projectTagStorage: projectTagStorage, tagStorage: tagStorage}
}

func (biz *createProjectBiz) CreateNewProject(ctx context.Context, data *model.ProjectCreation) error {
	//if err := data.Validate(); err != nil {
	//	return common.ErrValidation(err)
	//}

	if err := biz.store.CreateProject(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	// create tags
	if err := biz.tagStorage.CreateTagsByTagNames(ctx, int(tagModel.TypeProject), data.Tags); err != nil {
		return common.ErrCannotCreateEntity(tagModel.EntityName, err)
	}

	//// get tags id
	TagIds, err := biz.tagStorage.GetTagIdsByNames(ctx, data.Tags)
	if err != nil {
		return common.ErrCannotCreateEntity(tagModel.EntityName, err)
	}

	// update project tags
	if err := biz.projectTagStorage.UpdateProjectTagsByProjectId(ctx, data.Id, TagIds); err != nil {
		return common.ErrCannotCreateEntity(projectTagModel.EntityName, err)
	}

	return nil
}
