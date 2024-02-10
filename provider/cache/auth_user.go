package cache

import (
	"context"
	"fmt"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type AccountSto interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}
type authUserCache struct {
	store      AccountSto // mysql
	cacheStore Cache      // redis
}

func NewAuthUserCache(store AccountSto, cacheStore Cache) *authUserCache {
	return &authUserCache{store: store, cacheStore: cacheStore}
}

func (c *authUserCache) GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error) {
	var account *entities.Account

	key := "account:" + email // key store in redis

	err := c.cacheStore.Get(ctx, key, &account) // get data from redis
	if err != nil {
		fmt.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if account != nil {
		return account, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*12); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err

}
