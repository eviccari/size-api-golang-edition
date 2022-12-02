package repositories

import "github.com/eviccari/size-api-golang-edition/models"

type CreateRepository interface {
	Execute(models.SizeDTO) (err error)
}

type FindByIDRepository interface {
	Execute(id string) (dto models.SizeDTO, err error)
}

type UpdateRepository interface {
	Execute(dto models.SizeDTO) (updatedRowCount int64, err error)
}

type DeleteRepository interface {
	Execute(id string) (updatedRowCount int64, err error)
}
