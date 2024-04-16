package storage

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

func (s *sqlStore) CreateAssignee(ctx context.Context, data *model.TaskAssigneeCreation) error {

	if err := s.db.Create(data).Error; err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return common.NewCustomError(err, "duplicated task assignee", "ErrAssigneeDuplicated")
		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
			return common.ErrForeignKeyViolation(model.EntityName)
		} else {
			return common.ErrDB(err)
		}

		//if err != nil {
		//	return err
		//}
	}

	return nil
}
