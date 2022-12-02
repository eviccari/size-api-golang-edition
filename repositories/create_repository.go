package repositories

import (
	"context"
	"database/sql"

	"github.com/eviccari/size-api-golang-edition/models"
)

type CreateRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewCreateRepository(ctx context.Context, db *sql.DB) *CreateRepositoryImpl {
	return &CreateRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (cr *CreateRepositoryImpl) Execute(dto models.SizeDTO) (err error) {
	query := `insert into products.sizes (id, description, short_description, size_range)
	          values (?, ?, ?, ?)
	         `

	_, err = cr.db.ExecContext(cr.ctx, query, dto.ID, dto.Description, dto.ShortDescription, dto.SizeRange)
	return
}
