package constant

type Role int

const (
	UserRole Role = iota + 1
	VendorRole
	AdminRole
)

const (
	TypeVerifyEmail       = 1
	TypeResetPassword     = 2
	TypeProjectInvitation = 3
)
