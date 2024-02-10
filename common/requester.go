package common

type Requester interface {
	GetRole() int
	GetEmail() string
	GetID() int
}
