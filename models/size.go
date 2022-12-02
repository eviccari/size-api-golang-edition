package models

import (
	"github.com/eviccari/size-api-golang-edition/utils/stringutils"
	"github.com/google/uuid"
)

type Size struct {
	id               string
	description      string
	shortDescription string
	sizeRange        string
}

func NewSize(description, shortDescription, sizeRange string) Size {
	return Size{
		id:               uuid.NewString(),
		description:      description,
		shortDescription: shortDescription,
		sizeRange:        sizeRange,
	}
}

func ConvertForCreate(dto SizeDTO) Size {
	return NewSize(dto.Description, dto.ShortDescription, dto.SizeRange)
}

func ConvertForUpdate(dto SizeDTO) *Size {
	return &Size{
		id:               dto.ID,
		description:      dto.Description,
		shortDescription: dto.ShortDescription,
		sizeRange:        dto.SizeRange,
	}
}

func (s *Size) GetID() string {
	return s.id
}

func (s *Size) Validate() (errorList []string) {
	if stringutils.IsEmpty(s.id) {
		errorList = append(errorList, "id is required")
	}

	if stringutils.IsEmpty(s.description) {
		errorList = append(errorList, "description is required")
	}

	if stringutils.IsEmpty(s.shortDescription) {
		errorList = append(errorList, "shortDescription is required")
	}

	if stringutils.IsEmpty(s.sizeRange) {
		errorList = append(errorList, "sizeRange is required")
	}

	return errorList
}
