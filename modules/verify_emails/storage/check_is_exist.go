package verifyemailsstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *verifyEmailsStorage) CheckIsExist(ctx context.Context, email string, _type int) (*entities.VerifyEmail, error) {
	db := s.db
	data := &entities.VerifyEmail{}
	err := db.Table(entities.VerifyEmail{}.TableName()).Where("email = ? AND type = ?", email, _type).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
