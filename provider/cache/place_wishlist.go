package cache

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type PlaceWishListSto interface {
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error)
}

type placeWishListCache struct {
	store      PlaceWishListSto // mysql
	cacheStore Cache            // redis
}

func NewPlaceWishListCache(store PlaceWishListSto, cacheStore Cache) *placeWishListCache {
	return &placeWishListCache{store: store, cacheStore: cacheStore}
}

func (c *placeWishListCache) GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error) {
	var placeWishList []entities.PlaceWishList

	placeId := condition["place_id"].(int)
	userId := condition["user_id"].(int)

	for k, v := range condition {
		// parse interface{} to int
		if k == "place_id" {
			val := v.(int)
			placeId = val
		} else if k == "user_id" {
			val := v.(int)
			userId = val
		}
		// val := v.(int)
		// strVal := strconv.Itoa(val)
		// key += strVal // key store in redis
	}

	entityPlaceWishList := entities.PlaceWishList{
		PlaceId: placeId,
		UserId:  userId,
	}

	key := entityPlaceWishList.CacheKey()             // key store in redis
	err := c.cacheStore.Get(ctx, key, &placeWishList) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if placeWishList != nil {
		return placeWishList, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetByCondition(ctx, condition)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err

}
