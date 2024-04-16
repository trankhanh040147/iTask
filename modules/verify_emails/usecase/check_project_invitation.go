package verifyemailsusecase

import (
	"context"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/constant"
	"iTask/modules/project_members/model"
)

func (uc *verifyEmailsUseCase) CheckProjectInvitation(ctx context.Context, email string, code string, project_id int) (bool, error) {
	account, err := uc.accountStore.GetAccountByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, common.ErrEntityNotFound("account")
		}
		return false, err
	}

	// check if verify code, email and project_id is matching
	data, err := uc.verifyEmailsStore.GetProjectInvitation(ctx, email, code, project_id, constant.TypeProjectInvitation)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, common.ErrVerifyCodeIsNotMatching("verify code", nil)
		}
		return false, err
	}

	// check if verify code is expired
	if data.IsExpired() {
		return true, nil
	}

	// *check if user is already a member of project
	_, err = uc.projectMemberStore.GetMember(ctx, account.Id, project_id)
	if err == nil {
		return false, common.ErrEntityExisted("project member")
	}

	// *if everything is ok => add user to project as member
	projectMember := model.ProjectMemberCreation{
		ProjectId: project_id,
		UserId:    account.Id,
		Role:      model.RoleMember,
	}

	err = uc.projectMemberStore.CreateProjectMember(ctx, &projectMember)
	if err != nil {
		return false, common.ErrCannotUpdateEntity("account", err)
	}

	// todo: update verify code to expired/used
	data.IsUsed = 1
	err = uc.verifyEmailsStore.UpdateEntity(ctx, map[string]interface{}{"id": data.Id}, data)

	return false, nil
}
