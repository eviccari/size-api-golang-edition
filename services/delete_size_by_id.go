package services

import (
	"context"
	"errors"

	"github.com/eviccari/size-api-golang-edition/repositories"
	"github.com/eviccari/size-api-golang-edition/utils/stringutils"
)

type DeleteSizeByID struct {
	ctx  *context.Context
	repo repositories.DeleteRepository
}

func NewDeleteSizeByID(ctx *context.Context, repo repositories.DeleteRepository) *DeleteSizeByID {
	return &DeleteSizeByID{
		ctx:  ctx,
		repo: repo,
	}
}

func (dsByID *DeleteSizeByID) Execute(id string) (updatedRows int64, err error) {
	if stringutils.IsEmpty(id) {
		return 0, errors.New("id is required")
	}

	return dsByID.repo.Execute(id)
}
