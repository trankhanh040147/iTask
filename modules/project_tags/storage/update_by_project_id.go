package storage

import (
	"context"
	"iTask/modules/project_tags/model"
)

func (s *sqlStore) UpdateProjectTagsByProjectId(ctx context.Context, projectId int, tagIds []int) error {

	// insert or updates project tags by project id and tag ids
	projectTag := model.ProjectTag{}
	projectTag.ProjectId = projectId

	for _, tagId := range tagIds {
		projectTag.TagId = tagId
		projectTag.Id = 0

		// sol: using First and Update
		// if record existed, then update/ignore
		if err := s.db.Where("project_id = ? AND tag_id = ?", projectId, tagId).First(&projectTag).Error; err == nil {
			continue
		}

		// if record not existed, then insert
		if err := s.db.Create(&projectTag).Error; err != nil {
			return err
		}
	}
	return nil
}
