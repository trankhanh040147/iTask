package verifyemailsstorage

import (
	"context"
	"iTask/entities"
)

func (s *verifyEmailsStorage) Create(ctx context.Context, data *entities.VerifyEmail) (*entities.VerifyEmail, error) {
	db := s.db
	err := db.Table(entities.VerifyEmail{}.TableName()).Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
