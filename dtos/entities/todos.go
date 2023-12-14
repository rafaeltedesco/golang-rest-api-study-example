package dtos

import "time"

type TodoInputDTO struct {
	Title       string `json:"title"`
	PlannedDate string `json:"plannedDate"`
}

type TodoOutputDTO struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	PlannedDate time.Time `json:"plannedDate"`
	IsDone      bool      `json:"isDone"`
}
