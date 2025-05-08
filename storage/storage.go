package postgres

import (
	"database/sql"
)

type DB struct{
	DB *sql.DB
}

func Storage(db *sql.DB)*DB{
	return &DB{
		DB: db,
	}
}

