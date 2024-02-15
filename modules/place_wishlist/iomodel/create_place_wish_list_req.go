package iomodel

type CreatePlaceWishListReq struct {
	PlaceID    int `json:"place_id"`
	WishListID int `json:"wishlist_id"`
}
