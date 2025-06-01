package repository

import (
	"time"
	"github.com/nunusavi/task-manager/internal/model"

)


func CreateTask (task *model.Task) error{
	query := `INSERT INTO tasks (user_id, title, description, completed, created_at) VALUES ($1, $2, $3, $4, $5)`
	result, err := DB.Exec(query, task.UserID, task.Title, task.Description, task.Completed, time.Now())
	if err != nil{
		return err
	}
	id, _ := result.LastInsertId()
	task.ID = int(id)
	return nil
}
func GetTaskByUserID(userID int)([]model.Task, error){
		rows, err :=DB.Query(`SELECT id, user_id, title, description, completed, created_at FROM tasks WHERE user_id = $1`, userID) 
		if err!= nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, err 
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}