package constant

type Role int

const (
	UserRole Role = iota + 1
	VendorRole
	AdminRole
)

const (
	TypeVerifyEmail   = 1
	TypeResetPassword = 2
)

const (
	StatusActive   = 2
	StatusInactive = 1
)
