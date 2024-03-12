package model

import "context"

func (s *sqlStore) GetProjectTagsByProjectId(ctx context.Context, cond map[string]interface{}) (map[int]string, error) {
	data := make(map[int]string)

	//  select pt.project_id,
	//	group_concat(t.`name` separator ', ')
	//	from ProjectTags pt
	//	join Tags t on t.id = pt.tag_id
	//	group by project_id

	type result struct {
		ProjectId int
		Tags      string
	}

	var results []result

	if err := s.db.Table(ProjectTag{}.TableName()).
		Select("project_id, group_concat(t.`name` separator ', ') as tags").
		Joins("join Tags t on t.id = tag_id").
		Where(cond).
		Group("project_id").
		Scan(&results).Error; err != nil {

		return nil, err
	}

	for _, result := range results {
		data[result.ProjectId] = result.Tags
	}

	return data, nil
}
