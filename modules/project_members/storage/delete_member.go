package storage

import (
	"context"
	"iTask/modules/project_members/model"
)

func (s *sqlStore) DeleteMember(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Where(cond).Delete(model.ProjectMember{}).Error; err != nil {
		return err
	}

	return nil
}
