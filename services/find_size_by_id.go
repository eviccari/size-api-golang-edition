package services

import (
	"context"

	"github.com/eviccari/size-api-golang-edition/models"
	"github.com/eviccari/size-api-golang-edition/repositories"
)

type FindSizeByID struct {
	ctx  *context.Context
	repo repositories.FindByIDRepository
}

func NewFindSizeByID(ctx *context.Context, repo repositories.FindByIDRepository) *FindSizeByID {
	return &FindSizeByID{
		ctx:  ctx,
		repo: repo,
	}
}

func (fsbyID *FindSizeByID) Execute(id string) (dto models.SizeDTO, err error) {
	return fsbyID.repo.Execute(id)
}
