package repositories

import (
	"context"
	"database/sql"

	"github.com/eviccari/size-api-golang-edition/models"
)

type UpdateRepositoryImpl struct {
	ctx *context.Context
	db  *sql.DB
}

func NewUpdateRepositoryImpl(ctx *context.Context, db *sql.DB) *UpdateRepositoryImpl {
	return &UpdateRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (uri *UpdateRepositoryImpl) Execute(dto models.SizeDTO) (updatedRowCount int64, err error) {
	query := `
	          update products.sizes 
			     set description       = ?
				     short_description = ?
					 size_range        = ?
			   where id                = ? 	 
	         `
	res, err := uri.db.ExecContext(*uri.ctx, query, dto.Description, dto.ShortDescription, dto.SizeRange, dto.ID)
	if err != nil {
		return
	}

	updatedRowCount, err = res.RowsAffected()
	return
}
