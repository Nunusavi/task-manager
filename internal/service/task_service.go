package service

import(
	"time"

	"github.com/nunusavi/task-manager/internal/model"
	"github.com/nunusavi/task-manager/internal/repository"
)

func CreateTask(userID int, title, description string)(*model.Task, error){
task := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	err := repository.CreateTask(task)
	return task, err
}

func ListTasks(userID int) ([]model.Task, error) {
	return repository.GetTaskByUserID(userID)
}