package recipes

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/andresdb91/letmecook/pkg/commons"
)

type SQLiteComponentRepository struct {
	db       *sql.DB
	instance *SQLiteComponentRepository
}

func (rr *SQLiteComponentRepository) GetComponentByID(id string) (*recipes.Component, error) {
	componentItem := new(DBComponent)
	componentRows := []any{&componentItem.ID, &componentItem.Name}

	next := func() (*commons.QueryResult[DBComponent], error) {
		return GetQueryPage[DBComponent](rr.db, "SELECT * FROM recipes", componentItem, componentRows, 1, 0)
	}
	parser := func(dbComponent *[]*DBComponent) ([]*recipes.Component, error) {
		return parseDBComponent(dbComponent)
	}

	result, err := commons.ParseQueryPage[DBComponent, recipes.Component](next, parser)
	if err != nil {
		return nil, err
	} else {
		return result.Items[0], nil
	}
}

func parseDBComponent(dbComponent *[]*DBComponent) ([]*recipes.Component, error) {
	var components []*recipes.Component
	for _, dbComponent := range *dbComponent {
		component := recipes.Component{
			ID:   uuid.MustParse(dbComponent.ID),
			Name: dbComponent.Name,
		}
		components = append(components, &component)
	}
	return components, nil
}
