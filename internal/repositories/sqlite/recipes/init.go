package recipes

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
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

func Init(path string) map[string]interface{} {
	db := NewSQLiteDB(path)
	repositories := map[string]interface{}{
		"recipe": SQLiteRecipeRepository{db: db},
	}
	return repositories
}
