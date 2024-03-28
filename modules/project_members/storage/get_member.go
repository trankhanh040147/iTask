package storage

import (
	"context"
	"iTask/modules/project_members/model"
)

// GetMember will return a project member by user id and project id
func (s *sqlStore) GetMember(ctx context.Context, userId, projectId int) (*model.ProjectMember, error) {
	var member model.ProjectMember

	if err := s.db.
		Where("user_id = ? AND project_id = ?", userId, projectId).
		First(&member).Error; err != nil {
		return nil, err
	}

	return &member, nil
}
