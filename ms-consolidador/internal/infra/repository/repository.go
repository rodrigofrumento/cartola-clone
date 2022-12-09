package repository

import (
	"database/sql"

	"github.com/rodrigofrumento/cartola/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
