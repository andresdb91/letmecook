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

type DBComponent struct {
	ID   string
	Name string
	// Tags    []Tag
}

type DBComponentTag struct {
	ComponentID string
	TagID       string
}

type DBTag struct {
	ID   string
	Name string
}

type DBUnit struct {
	ID   string
	Name string
	// conversionTable map[string]float64
}

type DBUnitConversion struct {
	FromUnitID string
	ToUnitID   string
	Factor     float64
}
