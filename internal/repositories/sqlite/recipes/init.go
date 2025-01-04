package recipes

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var (
	RecipeRepository    SQLiteRecipeRepository
	ComponentRepository SQLiteComponentRepository
)

func NewSQLiteDB(path string) *sql.DB {
	var db *sql.DB
	if path == "" {
		db, _ = sql.Open("sqlite3", "file::memory:?cache=shared")
	} else {
		db, _ = sql.Open("sqlite3", "file:"+path+"?cache=shared")
	}
	return db
}

func Init(path string) {
	db := NewSQLiteDB(path)
	RecipeRepository = SQLiteRecipeRepository{db: db}
	ComponentRepository = SQLiteComponentRepository{db: db}
}
