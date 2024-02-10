package iomodel

type CreatePolicyReq struct {
	Data DataReqCreatePolicy `json:"data"`
}

type DataReqCreatePolicy struct {
	PlaceID    int          `json:"place_id"`
	ListPolicy []ListPolicy `json:"list_policy"`
}

type ListPolicy struct {
	GroupPolicyID int    `json:"group_policy_id"`
	Name          string `json:"name"`
}
