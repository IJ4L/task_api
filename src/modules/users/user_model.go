package users

import Task "taskManagement.com/src/modules/tasks"

type User struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment:true;index"`
	Name  string `json:"name" gorm:"type:varchar(30)"`
	Email string `json:"email" gorm:"type:varchar(20)"`
	Task  []Task.Task
}
