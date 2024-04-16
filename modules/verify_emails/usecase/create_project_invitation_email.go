package verifyemailsusecase

import (
	"context"
	"fmt"
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
	expiredTime := utils.GetExpiredTime(constant.ExpiredTimeProjectInvitation)

	// *create verify email as a record in table `verify_emails`
	record := &entities.VerifyEmail{
		Email:     email,
		Type:      constant.TypeProjectInvitation,
		ProjectId: project.Id,
		ScretCode: code,
		ExpiredAt: &expiredTime,
	}
	data, err := uc.verifyEmailsStore.Create(ctx, record)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("verify_emails", err)
	}

	//!logging
	fmt.Printf("project_id: %d, mail_id:%d\n", project.Id, data.Id)

	// todo: create record in ProjectMemberInvited
	//recordInvitation := projectMemberInvitedModel.ProjectMemberInvited{
	//	ProjectId:          project.Id,
	//	VerificationMailId: data.Id,
	//}
	//
	//err = uc.projectMemberInvitedStorage.CreateProjectMemberInvited(ctx, &recordInvitation)
	//if err != nil {
	//	return nil, common.ErrCannotCreateEntity(projectMemberInvitedModel.EntityName, err)
	//}

	return data, nil
}
