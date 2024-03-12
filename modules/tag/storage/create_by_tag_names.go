package storage

import (
	"context"
	"iTask/modules/tag/model"
	"strings"
)

func (s *sqlStore) CreateTagsByTagNames(ctx context.Context, tagType int, tags string) error {

	// split tag names
	tagNames := strings.Split(tags, ",")
	for _, tagName := range tagNames {
		tagName = strings.TrimSpace(tagName)

		tag := model.TagCreation{
			Name: tagName,
			Type: model.TagType(tagType),
		}

		// validate tag
		if err := tag.Validate(); err != nil {
			return err
		}

		// insert or updates tags by tag names
		// using First
		if err := s.db.Where("name = ? AND tag_type = ?", tagName, tagType).First(&tag).Error; err == nil {
			continue
		}

		if err := s.db.Create(&tag).Error; err != nil {
			return err
		}
	}

	return nil
}
