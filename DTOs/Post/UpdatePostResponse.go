package Post

import (
	"ToDoApp/Enums"
	"time"
)

type UpdatePostResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Name      string       `json:"name"`
	Status    Enums.Status `json:"status"`
	Time      time.Time    `json:"time"`
	UpdatedAt time.Time    `json:"updated_at"`
}
