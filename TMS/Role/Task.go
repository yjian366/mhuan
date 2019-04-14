package Component

import (
	"time"
)

type Task struct {
	ID int
	Creater 	User
	Owner		User
	BeginTime	time.Time
	EndTime 	time.Time
	FinishTime 	time.Time
	Content 	string
	Class 		int				//紧急程度,高，中，低。

	Grade		int				//分数，
	Comment 	string
	Status 		int				//任务状态：Finished,Delay,Running,Waiting

}
type TaskPlan struct{
	ID int
	Title string

}