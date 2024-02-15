package iomodel

type DeleteAmenityReq struct {
	IDPlace             int   `json:"place_id" form:"place_id"`
	ListConfigAmenityId []int `json:"list_config_amenity_id" form:"list_config_amenity_id"`
}
