package storage

import (
	"context"
	"iTask/modules/tag/model"
	"strings"
)

func (s *sqlStore) GetTagIdsByNames(ctx context.Context, tagNames string) ([]int, error) {
	tagNameLst := strings.Split(tagNames, ",")
	if len(tagNameLst) == 0 {
		return nil, nil
	}

	for i, tagName := range tagNameLst {
		tagNameLst[i] = strings.ToLower(strings.TrimSpace(tagName))
	}

	tag := model.Tag{}
	result := make([]int, len(tagNameLst))
	for i, tagName := range tagNameLst {
		tag.Name = tagName
		tag.Id = 0
		if err := s.db.Where("name = ?", tagName).First(&tag).Error; err != nil {
			continue
		}
		result[i] = tag.Id
	}

	return result, nil
}
