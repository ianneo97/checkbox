package requests

import "time"

type UpdateTasksRequestBody struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Status      string     `json:"status"`
}
