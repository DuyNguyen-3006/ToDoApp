package Enums

type Status string

const (
	StatusToDo       Status = "ToDo"
	StatusInProgress Status = "InProgress"
	StatusDone       Status = "Done"
)

func (status Status) IsValid() bool {
	switch status {
	case StatusToDo, StatusInProgress, StatusDone:
		return true
	}
	return false
}
