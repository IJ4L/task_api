package tasks

type UpdateTaskInputDto struct {
	Task string `json:"task" validate:"required"`
	Time string `json:"time" validate:"required"`
}
