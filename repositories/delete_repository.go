package repositories

import (
	"context"
	"database/sql"
)

type DeleteRepositoryImpl struct {
	ctx *context.Context
	db  *sql.DB
}

func NewDeleteRepositoryImpl(ctx *context.Context, db *sql.DB) *DeleteRepositoryImpl {
	return &DeleteRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (dri *DeleteRepositoryImpl) Execute(id string) (updatedRowCount int64, err error) {
	query := "delete from products.sizes where id = ?"
	res, err := dri.db.ExecContext(*dri.ctx, query, id)
	if err != nil {
		return
	}

	return res.RowsAffected()
}
