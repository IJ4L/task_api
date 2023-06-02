package tasks

type Task struct {
	ID     int64  `json:"id" gorm:"primary_key;auto_increment:true;index"`
	Task   string `json:"name_task" gorm:"type:varchar(255)"`
	Time   string `json:"time" gorm:"type:varchar(40)"`
	UserID uint   `json:"user_id"`
}
