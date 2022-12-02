package repositories

import (
	"context"
	"database/sql"

	"github.com/eviccari/size-api-golang-edition/models"
)

type FindByIDRepositoryImpl struct {
	ctx *context.Context
	db  *sql.DB
}

func NewFindByIDRepositoryImpl(ctx *context.Context, db *sql.DB) *FindByIDRepositoryImpl {
	return &FindByIDRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (fbID *FindByIDRepositoryImpl) Execute(id string) (dto models.SizeDTO, err error) {
	query := `
	           select id, 
			          description, 
					  short_description as \"shortDescription\", 
					  size_range as \"sizeRange\"
				from products.sizes
			   where id = ?		  
			`
	err = fbID.db.QueryRowContext(*fbID.ctx, query, id).Scan(dto)
	return
}
