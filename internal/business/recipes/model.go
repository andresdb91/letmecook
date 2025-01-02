package recipes

import "github.com/google/uuid"

type Tag struct {
	ID      uuid.UUID
	Name    string
	Indexed bool
}

type Component struct {
	ID      uuid.UUID
	Name    string
	Tags    []Tag
	Recipes []*Recipe
}

type Unit struct {
	ID              uuid.UUID
	Name            string
	conversionTable map[*Unit]float64
}

type Quantity struct {
	Amount float64
	Unit   Unit
}

func (q *Quantity) ConvertTo(unit *Unit) {
	q.Amount = q.Amount * q.Unit.conversionTable[unit]
	q.Unit = *unit
}

type Ingredient struct {
	Component    *Component
	Amount       Quantity
	Replacements []*Component
}

type Step struct {
	Description string
	Components  []*Component
	Picture     string
}

type Recipe struct {
	ID                 uuid.UUID
	ResultingComponent *Component
	Name               string
	Ingredients        []Ingredient
	Steps              []Step
	Tags               []Tag
}
