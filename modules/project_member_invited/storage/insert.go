package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/project_member_invited/model"
)

func (s *sqlStore) CreateProjectMemberInvited(ctx context.Context, data *model.ProjectMemberInvited) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
