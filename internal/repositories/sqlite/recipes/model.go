package recipes

type DBRecipe struct {
	ID                   string
	Name                 string
	ResultingComponentID string
	// Ingredients        []string
	// Steps              []string
	// Tags               []string
}

type DBRecipeIngredient struct {
	ID          string
	RecipeID    string
	ComponentID string
	Amount      float64
	UnitID      string
}

type DBRecipeStep struct {
	ID          string
	RecipeID    string
	Description string
	// Components  []string
	Picture string
}

type DBRecipeStepComponent struct {
	StepID      string
	ComponentID string
}

type DBRecipeTags struct {
	RecipeID string
	TagID    string
}

type DBRecipeTag struct {
	ID   string
	Name string
}

type Component struct {
	ID   string
	Name string
	// Tags    []Tag
}

type ComponentTag struct {
	ComponentID string
	TagID       string
}

type Tag struct {
	ID   string
	Name string
}

type Unit struct {
	ID   string
	Name string
	// conversionTable map[string]float64
}

type UnitConversion struct {
	FromUnitID string
	ToUnitID   string
	Factor     float64
}
