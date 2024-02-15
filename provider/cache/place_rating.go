package cache

import (
	"context"
	"log"
	"iTask/common"
	"iTask/entities"
	"time"
)

type PlaceSto interface {
	GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error)
}

type placeStoCache struct {
	store      PlaceSto // mysql
	cacheStore Cache    // redis
}

func NewPlaceStoCache(store PlaceSto, cacheStore Cache) *placeStoCache {
	return &placeStoCache{store: store, cacheStore: cacheStore}
}

func (c *placeStoCache) GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error) {
	place := entities.Place{}
	place.Id = int(placeId)

	key := place.CacheKeyPlaceRating()
	var ratingAverage *float64

	err := c.cacheStore.Get(ctx, key, &ratingAverage) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if ratingAverage != nil {
		return ratingAverage, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetRatingAverageByPlaceId(ctx, placeId)
	if err != nil {
		return nil, err
	}

	if u == nil {
		defaulRating := 0.0
		u = &defaulRating
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
