package services

import (
	"context"

	"github.com/eviccari/size-api-golang-edition/models"
	"github.com/eviccari/size-api-golang-edition/repositories"
	"github.com/eviccari/size-api-golang-edition/utils/stringutils"
)

type CreateSize struct {
	ctx  context.Context
	repo repositories.CreateRepository
}

func NewCreateSize(ctx context.Context, repo repositories.CreateRepository) *CreateSize {
	return &CreateSize{
		ctx:  ctx,
		repo: repo,
	}
}

func (cns *CreateSize) Execute(dto models.SizeDTO) (err error) {
	model := models.ConvertForCreate(dto)
	msgs := model.Validate()

	if len(msgs) > 0 {
		return stringutils.BuildErrorFromMessages(msgs)
	}

	dto.ID = model.GetID()
	return cns.repo.Execute(dto)
}
