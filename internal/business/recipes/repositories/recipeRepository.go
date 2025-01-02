package recipes

import (
	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/google/uuid"
)

type ComponentRepository interface {
	GetComponentByID(id string) (*recipes.Component, error)
	GetComponentsByTags(tags []string) ([]*recipes.Component, error)
}

type UnitRepository interface {
}

type QuantityRepository interface {
}

type IngredientRepository interface {
}

type StepRepository interface {
}

type RecipeRepository interface {
	GetRecipeByID(id string) (*recipes.Recipe, error)
	GetRecipesByComponents(components []string) ([]*recipes.Recipe, error)
	GetRecipesByKeywords(keywords []string) ([]*recipes.Recipe, error)
	GetRecipesByName(name string) ([]*recipes.Recipe, error)
	GetRecipesByTags(tags []string) ([]*recipes.Recipe, error)
	GetTopRatedRecipes() ([]*recipes.Recipe, error)
	GetTopSearchedRecipes() ([]*recipes.Recipe, error)
	GetNewestRecipes() ([]*recipes.Recipe, error)
	GetRandomRecipe() (*recipes.Recipe, error)
}

type RecipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) *RecipeService {
	return &RecipeService{repository: repository}
}

func (rs *RecipeService) GetRecipeByID(id uuid.UUID) (*recipes.Recipe, error) {
	return rs.repository.GetRecipeByID(id.String())
}

func (rs *RecipeService) GetRecipesByComponents(components []*recipes.Component) ([]*recipes.Recipe, error) {
	var componentIDs []string
	for _, component := range components {
		componentIDs = append(componentIDs, component.ID.String())
	}
	return rs.repository.GetRecipesByComponents(componentIDs)
}

func (rs *RecipeService) GetRecipesByKeywords(keywords []string) ([]*recipes.Recipe, error) {
	return rs.repository.GetRecipesByKeywords(keywords)
}

func (rs *RecipeService) GetRecipesByName(name string) ([]*recipes.Recipe, error) {
	return rs.repository.GetRecipesByName(name)
}

func (rs *RecipeService) GetRecipesByTags(tags []*recipes.Tag) ([]*recipes.Recipe, error) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID.String())
	}
	return rs.repository.GetRecipesByTags(tagIDs)
}

func (rs *RecipeService) GetTopRatedRecipes() ([]*recipes.Recipe, error) {
	return rs.repository.GetTopRatedRecipes()
}

func (rs *RecipeService) GetTopSearchedRecipes() ([]*recipes.Recipe, error) {
	return rs.repository.GetTopSearchedRecipes()
}

func (rs *RecipeService) GetNewestRecipes() ([]*recipes.Recipe, error) {
	return rs.repository.GetNewestRecipes()
}

func (rs *RecipeService) GetRandomRecipe() (*recipes.Recipe, error) {
	return rs.repository.GetRandomRecipe()
}
