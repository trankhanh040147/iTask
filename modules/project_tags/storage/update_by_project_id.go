package storage

import (
	"context"
	"iTask/modules/project_tags/model"
)

func (s *sqlStore) UpdateProjectTagsByProjectId(ctx context.Context, projectId int, tagIds []int) error {

	// insert or updates project tags by project id and tag ids
	for _, tagId := range tagIds {
		projectTag := model.ProjectTag{
			ProjectId: projectId,
			TagId:     tagId,
		}

		// method 1: using Save
		if err := s.db.Save(&projectTag).Error; err != nil {
			return err
		}

		// method 2: using First and Update
		//if err := s.DB.Where("project_id = ? AND tag_id = ?", projectId, tagId).First(&projectTag).Error; err != nil {
		//	if err := s.DB.Create(&projectTag).Error; err != nil {
		//		return err
		//	}
		//} else {
		//	if err := s.DB.Model(&projectTag).Update("deleted_at", nil).Error; err != nil {
		//		return err
		//	}
		//}
	}
	return nil
}
