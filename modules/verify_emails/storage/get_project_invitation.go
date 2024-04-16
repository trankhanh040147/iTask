package verifyemailsstorage

import (
	"context"
	"iTask/entities"
)

func (s *verifyEmailsStorage) GetProjectInvitation(ctx context.Context, email string, verifyCode string, project_id int, _type int) (*entities.VerifyEmail, error) {
	db := s.db
	data := &entities.VerifyEmail{}
	err := db.Table(entities.VerifyEmail{}.TableName()).Where("email = ? and scret_code = ? and project_id = ? and type = ?", email, verifyCode, project_id, _type).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
