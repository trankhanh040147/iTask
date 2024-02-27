package common

type Requester interface {
	GetRole() string
	GetEmail() string
	GetUserId() int
}

const CurrentUser = "current_user"

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin"
}
