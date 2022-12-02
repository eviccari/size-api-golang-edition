package models

type CreateSize interface {
	Execute(dto SizeDTO) (err error)
}

type FindSizeByID interface {
	Execute(id string) (dto SizeDTO, err error)
}

type UpdateSize interface {
	Execute(dto SizeDTO) (updatedRows int64, err error)
}

type DeleteSize interface {
	Execute(id string) (updatedRows int64, err error)
}
