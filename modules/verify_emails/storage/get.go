package verifyemailsstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *verifyEmailsStorage) Get(ctx context.Context, email string, verifyCode string, _type int) (*entities.VerifyEmail, error) {
	db := s.db
	data := &entities.VerifyEmail{}
	err := db.Table(entities.VerifyEmail{}.TableName()).Where("email = ? AND scret_code = ? and type = ?", email, verifyCode, _type).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
