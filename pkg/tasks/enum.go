package tasks

type TaskStatus int

const (
	NOT_URGENT = iota
	DUE_SOON
	OVERDUE
)

func (ts TaskStatus) String() string {
	return []string{"not_urgent", "due_soon", "overdue"}[ts]
}
