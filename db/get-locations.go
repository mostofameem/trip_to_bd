package db

import (
	"context"
	"database/sql"
	"log/slog"
	"post-service/web/utils"

	sq "github.com/Masterminds/squirrel"
)

func (repo *LocationTypeRepo) GetLocations(ctx context.Context, params utils.PaginationParams) (*[]Location, error) {
	limit, offset := ConfigPageSize(params.Page, params.Limit)

	Query := NewQueryBuilder().
		Select("*").
		From(repo.table)
	for k, v := range params.Filters {
		Query = Query.Where(sq.Eq{k: v})
	}
	Query = Query.Limit(uint64(limit)).
		Offset(uint64(offset))
	query, args, err := Query.ToSql()
	if err != nil {
		slog.Error("faild to build query")
		return nil, err
	}

	var locations *[]Location
	err = repo.db.QueryRowContext(ctx, query, args...).Scan(&locations)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("failed to execute query")
		return nil, err
	}

	return locations, nil
}

func ConfigPageSize(page, limit int) (int, int) {
	PageLimit := 20
	Offset := 0

	PageLimit = min(limit, PageLimit)
	Offset = PageLimit * page

	return PageLimit, Offset
}
