package Post

import (
	"ToDoApp/Enums"
)

type PostRequest struct {
	Title  string       `json:"title" validate:"required"`
	Body   string       `json:"body" validate:"required"`
	Name   string       `json:"name" validate:"required"`
	Status Enums.Status `json:"status" validate:"required"`
}
