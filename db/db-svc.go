package db

import (
	"database/sql"
	"post-service/config"
)

func NewDB(dbCnf *config.DBConfig) *sql.DB {
	db := connect(dbCnf)
	return db
}
