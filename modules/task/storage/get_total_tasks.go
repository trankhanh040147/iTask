package storage

//func (s *sqlStore) GetTotalTasks(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
//	var data *model.TodoItem
//
//	if err := s.db.Where(cond).First(&data).Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return nil, common.RecordNotFound
//		}
//
//		return nil, common.ErrDB(err)
//	}
//
//	return data, nil
//}
