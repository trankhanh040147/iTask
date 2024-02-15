package wishlistiomodel

type CreateWishListReq struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
