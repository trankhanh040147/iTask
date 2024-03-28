package verifyemailsusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"iTask/modules/project/model"
	"iTask/utils"
)

func (uc *verifyEmailsUseCase) CreateProjectInvitationEmail(ctx context.Context, email string, project *model.Project) (*entities.VerifyEmail, error) {
	// create verify code
	code := utils.GenerateRandomCode(constant.LengthRandomCode)
	//code := "pid" + string(project)

	// set expired time
	expiredTime := utils.GetExpiredTime(constant.ExpiredTimeVerifyEmail)

	// create verify email
	record := &entities.VerifyEmail{
		Email:     email,
		Type:      constant.TypeVerifyEmail,
		ScretCode: code,
		ExpiredAt: &expiredTime,
	}
	data, err := uc.verifyEmailsStore.Create(ctx, record)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("verify_emails", err)
	}

	// todo: create record in ProjectMemberInvited
	
	return data, nil
}
