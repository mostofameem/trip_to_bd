package db

import (
	"database/sql"
	"post-service/config"
	"time"
)

type LocationTypeRepo struct {
	db    *sql.DB
	table string
}

func NewLocationTypeRepo(cnf *config.DBConfig) *LocationTypeRepo {
	db := NewDB(cnf)
	return &LocationTypeRepo{
		db:    db,
		table: "locations",
	}
}

type Location struct {
	ID         int       `db:"id"`
	Title      string    `db:"title" json:"title" validate:"required"`
	BestTime   string    `db:"best_time" json:"best_time" validate:"required"`
	PictureUrl string    `db:"picture_url" json:"picture_url" validate:"required"`
	Rating     float32   `db:"rating" json:"rating"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Isactive   bool      `db:"is_active"`
}
