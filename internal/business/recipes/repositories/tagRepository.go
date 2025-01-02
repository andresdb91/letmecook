package recipes

import (
	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
)

type TagRepository interface {
	GetTagByID(id string) (*recipes.Tag, error)
	GetTagsByIDList(ids []string) ([]*recipes.Tag, error)
}

type TagService struct {
	repository TagRepository
}

func NewTagService(repository TagRepository) *TagService {
	return &TagService{repository: repository}
}

func (ts *TagService) GetTagByID(id string) (*recipes.Tag, error) {
	return ts.repository.GetTagByID(id)
}

func (ts *TagService) GetIndexedTags(tags []*recipes.Tag) ([]*recipes.Tag, error) {
	var indexedTagIDs []string
	for _, tag := range tags {
		if tag.Indexed {
			indexedTagIDs = append(indexedTagIDs, tag.ID.String())
		}
	}
	return ts.repository.GetTagsByIDList(indexedTagIDs)
}
