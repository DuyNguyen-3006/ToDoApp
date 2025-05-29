package Post

import (
	"ToDoApp/Enums"
	"time"
)

type PostResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Name      string       `json:"name"`
	Status    Enums.Status `json:"status"`
	Time      time.Time    `json:"time"`
	CreatedAt time.Time    `json:"created_at"`
}
