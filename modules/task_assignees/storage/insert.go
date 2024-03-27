package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

func (s *sqlStore) CreateAssignee(ctx context.Context, data *model.TaskAssigneeCreation) error {

	if err := s.db.Create(data).Error; err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.New(fmt.Sprintf("duplicated task assignee for: task_id: %d, user_id: %d", data.TaskId, data.UserId))
		} else {
			return common.ErrDB(err)
		}
	}

	return nil
}
