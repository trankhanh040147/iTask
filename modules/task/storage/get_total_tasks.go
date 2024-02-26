package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task/model"
)

func (s *sqlStore) GetTotalTasks(ctx context.Context, cond map[string]interface{}) (map[int]int, error) {
	data := make(map[int]int)

	// Select project_id, count(*) as total_tasks from Tasks group by project_id

	type result struct {
		ProjectID  int
		TotalTasks int
	}

	var results []result

	if err := s.db.Table(model.Task{}.TableName()).
		//if err := s.db.Model(&model.Task{}).
		Select("project_id, count(id) as total_tasks").
		Where(cond).
		Group("project_id").
		//Pluck("project_id, total_tasks", &data).Error; err != nil {
		Scan(&results).Error; err != nil {

		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)

	}

	for _, result := range results {
		data[result.ProjectID] = result.TotalTasks
	}

	return data, nil
}
