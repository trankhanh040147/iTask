package accountstorage

import (
	"context"
	"fmt"
)

// ExecTx executes a function within a database transaction
func (s *accountStorage) execTx(ctx context.Context, fn func(*accountStorage) error) error {
	db := s.db.Begin()

	err := fn(&accountStorage{db: db})
	if err != nil {
		if rbErr := db.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	if err = db.Commit().Error; err != nil {
		return fmt.Errorf("commit err: %v", err)
	}
	return nil
}
