package verifyemailsstorage

import (
	"context"
	"iTask/entities"
)

func (s *verifyEmailsStorage) UpdateEntity(ctx context.Context, cond map[string]interface{}, data *entities.VerifyEmail) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
