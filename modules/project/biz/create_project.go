package biz

import (
	"context"
	"fmt"
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

type ProjectMemberStorage interface {
	CreateProjectMember(ctx context.Context, projectId int, userId int) error
}

type createProjectBiz struct {
	store                CreateProjectStorage
	projectTagStorage    ProjectTagStorage
	tagStorage           TagStorage
	projectMemberStorage ProjectMemberStorage
	requester            common.Requester
}

func NewCreateProjectBiz(store CreateProjectStorage, projectTagStorage ProjectTagStorage, tagStorage TagStorage, projectMemberStorage ProjectMemberStorage, requester common.Requester) *createProjectBiz {
	return &createProjectBiz{
		store:                store,
		projectTagStorage:    projectTagStorage,
		tagStorage:           tagStorage,
		projectMemberStorage: projectMemberStorage,
		requester:            requester,
	}
}

func (biz *createProjectBiz) CreateNewProject(ctx context.Context, data *model.ProjectCreation) error {
	//if err := data.Validate(); err != nil {
	//	return common.ErrValidation(err)
	//}

	// todo: testing CreateBy
	data.CreatedBy = biz.requester.GetUserId()

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

	// !logging
	fmt.Printf("<____________> Project w/ id %d has been created by user w/ id %d\n", data.Id, data.CreatedBy)

	// todo: add project_members
	if err := biz.projectMemberStorage.CreateProjectMember(ctx, data.Id, data.CreatedBy); err != nil {
		return common.ErrCannotCreateEntity("project_members", err)
	}
	// todo: implement: worker(redis) | pubsub...

	return nil
}
