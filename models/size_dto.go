package models

import "encoding/json"

type SizeDTO struct {
	ID               string `json:"id"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	SizeRange        string `json:"sizeRange"`
}

func NewSizeDTO(description, shortDescription, sizeRange string) *SizeDTO {
	return &SizeDTO{
		Description:      description,
		ShortDescription: shortDescription,
		SizeRange:        sizeRange,
	}
}

func (s *SizeDTO) ToJSONString() string {
	j, _ := json.Marshal(s)
	return string(j)
}
