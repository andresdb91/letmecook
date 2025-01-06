package recipes

import (
	"github.com/google/uuid"

	"github.com/andresdb91/letmecook/pkg/commons"
)

type Tag struct {
	ID      uuid.UUID
	Name    string
	Indexed bool
}

type Component struct {
	ID      uuid.UUID
	Name    string
	Tags    *commons.PagedList[Tag]
	Recipes *commons.PagedList[Recipe]
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
	Replacements *commons.PagedList[Component]
}

type Step struct {
	Description string
	Components  *commons.PagedList[Component]
	Picture     string
}

type Recipe struct {
	ID                 uuid.UUID
	Name               string
	ResultingComponent *Component
	Ingredients        *commons.PagedList[Ingredient]
	Steps              *commons.PagedList[Step]
	Tags               *commons.PagedList[Tag]
}
