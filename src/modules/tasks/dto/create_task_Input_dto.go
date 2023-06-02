package tasks

type CreateTaskInputDto struct {
	Task   string `json:"task" validate:"required"`
	Time   string `json:"time" validate:"required"`
	UserID uint   `json:"user_id" validate:"required"`
}
