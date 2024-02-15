package bookingdetailstorage

import (
	"context"
	"fmt"
)

// ExecTx executes a function within a database transaction
func (s *bookingDetailStorage) execTx(ctx context.Context, fn func(*bookingDetailStorage) error) error {
	db := s.db.Begin()

	err := fn(&bookingDetailStorage{db: db})
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
