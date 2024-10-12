package utils

import (
	"math"
	"net/http"
	"strconv"
)

type FilterParams map[string]any
type PaginationParams struct {
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
	Filter    FilterParams
}

const (
	maxLimit     = 100.0
	pageKey      = "page"
	limitKey     = "limit"
	sortByKey    = "sortBy"
	sortOrderKey = "sortOrder"
)

func parsePage(r *http.Request) int {
	pageStr := r.URL.Query().Get(pageKey)
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	page = int64(math.Max(0.0, float64(page)))
	return int(page)
}

func parseLimit(r *http.Request) int {
	limitStr := r.URL.Query().Get(limitKey)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	limit = int64(math.Max(0.0, math.Min(maxLimit, float64(limit))))
	return int(limit)
}

func countTotalPages(limit, totalItems int) int {
	return int(math.Ceil(float64(totalItems) / math.Max(1.0, float64(limit))))
}

// func GetPaginationParams(r *http.Request , defaultSortBy, defaultSortOrder string)PaginationParams{
// 	params:=PaginationParams{
// 		Page: 0,
// 		Limit: 10,
// 		SortBy: defaultSortBy,
// 		SortOrder: defaultSortOrder,

// 	}
// }
