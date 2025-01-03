package recipes

import (
	"github.com/google/uuid"

	"github.com/andresdb91/letmecook/pkg/commons"
)

type ServiceInterface interface {
	GetRepository() interface{}
}

type RecipeService struct {
	repository RecipeRepository
}

func (rs *RecipeService) GetRepository() interface{} {
	return rs.repository
}

func (rs *RecipeService) GetRecipeByID(id uuid.UUID) (*Recipe, error) {
	return rs.repository.GetRecipeByID(id.String())
}

func (rs *RecipeService) GetRecipesByComponents(components []*Component) ([]*Recipe, error) {
	var componentIDs []string
	for _, component := range components {
		componentIDs = append(componentIDs, component.ID.String())
	}
	return rs.repository.GetRecipesByComponents(componentIDs)
}

func (rs *RecipeService) GetRecipesByKeywords(keywords []string) ([]*Recipe, error) {
	return rs.repository.GetRecipesByKeywords(keywords)
}

func (rs *RecipeService) GetRecipesByName(name string) ([]*Recipe, error) {
	return rs.repository.GetRecipesByName(name)
}

func (rs *RecipeService) GetRecipesByTags(tags []*Tag) ([]*Recipe, error) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID.String())
	}
	return rs.repository.GetRecipesByTags(tagIDs)
}

func (rs *RecipeService) GetTopRatedRecipes() ([]*Recipe, error) {
	return rs.repository.GetTopRatedRecipes()
}

func (rs *RecipeService) GetTopSearchedRecipes() ([]*Recipe, error) {
	return rs.repository.GetTopSearchedRecipes()
}

func (rs *RecipeService) GetNewestRecipes() ([]*Recipe, error) {
	return rs.repository.GetNewestRecipes()
}

func (rs *RecipeService) GetRandomRecipe() (*Recipe, error) {
	return rs.repository.GetRandomRecipe()
}

func (rs *RecipeService) GetAllRecipes() (*commons.PagedList[Recipe], error) {
	return rs.repository.GetAllRecipes()
}
