package verifyemailsstorage

import (
	"context"
	"iTask/entities"
)

func (s *verifyEmailsStorage) FindProjectInvitation(ctx context.Context, cond map[string]interface{}) ([]entities.VerifyEmail, error) {
	var emails []entities.VerifyEmail

	db := s.db.Table(entities.VerifyEmail{}.TableName())

	err := db.Where(cond).Find(&emails).Order("expired_at desc").Error
	if err != nil {
		return nil, err
	}

	return emails, nil
}
