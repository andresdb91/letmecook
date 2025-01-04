package recipes

import (
	"database/sql"

	"github.com/andresdb91/letmecook/pkg/commons"
)

func GetQueryPage[T any](db *sql.DB, query string, scanItem *T, scanRows []any, pageSize int, pageNumber int) (*commons.QueryResult[T], error) {
	query = query + " LIMIT ? OFFSET ?"
	rows, err := db.Query(query, pageSize+1, pageNumber*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultList []*T
	// var pagedRecipes commons.PagedList[T]
	var scanError error
	for rows.Next() {
		if scanError = rows.Scan(scanRows...); scanError != nil {
			// return &pagedRecipes, err
			break
		}
		resultList = append(resultList, scanItem)
	}

	var next func() (*commons.QueryResult[T], error)
	if len(resultList) > pageSize {
		next = func() (*commons.QueryResult[T], error) {
			return GetQueryPage[T](db, query, scanItem, scanRows, pageSize, pageNumber+1)
		}
	} else {
		next = nil
	}

	result := commons.QueryResult[T]{
		Items:       resultList,
		Next:        next,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}

	if scanError != nil {
		return &result, scanError
	} else {
		return &result, nil
	}
}
