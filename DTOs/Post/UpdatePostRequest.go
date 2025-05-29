package Post

import "ToDoApp/Enums"

type UpdatePostRequest struct {
	Title  string       `json:"title"`
	Body   string       `json:"body"`
	Name   string       `json:"name"`
	Status Enums.Status `json:"status"`
}
