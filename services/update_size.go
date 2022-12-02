package services

import (
	"context"

	"github.com/eviccari/size-api-golang-edition/models"
	"github.com/eviccari/size-api-golang-edition/repositories"
	"github.com/eviccari/size-api-golang-edition/utils/stringutils"
)

type UpdateSize struct {
	ctx  *context.Context
	repo repositories.UpdateRepository
}

func NewUpdateSizeData(ctx *context.Context, repo repositories.UpdateRepository) *UpdateSize {
	return &UpdateSize{
		ctx:  ctx,
		repo: repo,
	}
}

func (usd *UpdateSize) Execute(dto models.SizeDTO) (updatedRows int64, err error) {
	model := models.ConvertForUpdate(dto)
	msgs := model.Validate()

	if len(msgs) > 0 {
		err = stringutils.BuildErrorFromMessages(msgs)
		return
	}

	return usd.repo.Execute(dto)
}
