package storage

import (
	"context"
	"iTask/modules/project_members/model"
)

func (s *sqlStore) UpdateMemberRole(ctx context.Context, projectId, memberId, role int) error {
	if err := s.db.Model(&model.ProjectMember{}).
		Where("project_id = ? and user_id = ?", projectId, memberId).
		Update("role", role).Error; err != nil {
		return err
	}

	return nil
}
