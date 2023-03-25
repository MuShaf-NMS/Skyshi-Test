package dto

type CreateTodo struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required"`
	IsActive        *bool  `json:"is_active" validate:"omitempty"`
}

type UpdateTodo struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
	Status   string `json:"status"`
}
