package verifyemailsstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *verifyEmailsStorage) Update(ctx context.Context, email string, _type int, verifyEmail *entities.VerifyEmail) error {
	db := s.db
	data := entities.VerifyEmail{}
	if err := db.Table(data.TableName()).Where("email = ? and type = ?", email, _type).Updates(verifyEmail).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
