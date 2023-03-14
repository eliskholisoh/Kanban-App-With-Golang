package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("id").(string)
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	idTask := r.URL.Query().Get("task_id")
	if idTask == "" {
		idUsr, err := strconv.Atoi(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		task, err := t.taskService.GetTasks(r.Context(), idUsr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)

	} else {

		tskID, err := strconv.Atoi(idTask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		task, err := t.taskService.GetTaskByID(r.Context(), tskID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
	// TODO: answer here
}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	// TODO: answer here
	if task.Title == "" || task.Description == "" || task.CategoryID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}
	idTsk1 := r.Context().Value("id")
	if idTsk1 == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	idTsk := r.Context().Value("id").(string)
	usrID, err := strconv.Atoi(idTsk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	etyTask := entity.Task{Title: task.Title, Description: task.Description, CategoryID: task.CategoryID, UserID: usrID}
	tasks, err := t.taskService.StoreTask(r.Context(), &etyTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": usrID, "task_id": tasks.ID, "message": "success create new task"})
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	idTsk := r.URL.Query().Get("task_id")
	tskID, err := strconv.Atoi(idTsk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	err = t.taskService.DeleteTask(r.Context(), tskID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	usrID := r.Context().Value("id")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": usrID, "task_id": tskID, "message": "success delete task"})
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	// TODO: answer here
	tskId1 := r.Context().Value("id")
	if tskId1 == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	tskId := r.Context().Value("id").(string)
	usrID, err := strconv.Atoi(tskId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	etyTask := entity.Task{ID: task.ID, Title: task.Title, Description: task.Description, UserID: usrID}
	tasks, err := t.taskService.UpdateTask(r.Context(), &etyTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": tskId, "task_id": tasks.ID, "message": "success update task"})
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
