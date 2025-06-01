package handler

import(
	"encoding/json"
	"net/http"

	"github.com/nunusavi/task-manager/internal/middleware"
	"github.com/nunusavi/task-manager/internal/service"
)

type CreatetaskRequest struct{
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	userID, ok := middleware.GetUserIDFromContext(r)
	if !ok{
		http.Error(w,"unauthorized", http.StatusUnauthorized)
		return
	}
	var req CreatetaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err !=nil{
		http.Error(w,"Invalid request", http.StatusBadRequest)
	return
	}

	task, err := service.CreateTask(userID, req.Title, req.Description)
	if err!=nil{
		http.Error(w,"Failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request){
	userID, ok := middleware.GetUserIDFromContext(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	tasks, err := service.ListTasks(userID)
	if err != nil{
		http.Error(w, "Failed to list tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}