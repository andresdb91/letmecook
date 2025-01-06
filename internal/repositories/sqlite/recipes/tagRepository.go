package recipes

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/andresdb91/letmecook/pkg/commons"
)

type SQLiteTagRepository struct {
	db *sql.DB
}

func (r SQLiteTagRepository) GetTagsByOwnerID(ownerID string) (*commons.PagedList[recipes.Tag], error) {
	pageSize := 10
	tagItem := new(DBTag)
	tagRows := []any{&tagItem.ID, &tagItem.Name}

	next := func() (*commons.QueryResult[DBTag], error) {
		return GetQueryPage[DBTag](r.db, "SELECT * FROM tags WHERE owner_id = ?", tagItem, tagRows, pageSize, ownerID)
	}

	return GetPagedTags(next)
}
