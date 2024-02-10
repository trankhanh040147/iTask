package policiesstorage

import "gorm.io/gorm"

type policyStorage struct {
	db *gorm.DB
}

func NewPolicyStorage(db *gorm.DB) *policyStorage {
	return &policyStorage{db: db}
}
