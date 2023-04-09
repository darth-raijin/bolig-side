package dtos

import (
	"github.com/darth-raijin/bolig-side/api/models/enums"
	"github.com/google/uuid"
)

type Question struct {
	ID       uuid.UUID    `json:"id" format:"uuid"`
	Question string       `json:"question"`
	Rating   enums.Rating `json:"rating"`
}
