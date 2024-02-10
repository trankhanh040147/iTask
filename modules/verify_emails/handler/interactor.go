package verifyemailshanlder

import (
	"context"
)

type verifyEmailsUseCase interface {
	CheckVerifyCodeIsMatching(ctx context.Context, email string, code string) (bool, error)
	CheckResetCodePasswordIsMatching(ctx context.Context, email string, code string) error
}

type verifyEmailsHandler struct {
	verifyEmailsUC verifyEmailsUseCase
}

func NewVerifyEmailsHandler(verifyEmailsUC verifyEmailsUseCase) *verifyEmailsHandler {
	return &verifyEmailsHandler{verifyEmailsUC: verifyEmailsUC}
}
