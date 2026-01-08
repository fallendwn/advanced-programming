package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DenisLi/Assignment2/internal/model"
	"github.com/DenisLi/Assignment2/internal/queue"
	"github.com/DenisLi/Assignment2/internal/store"
)

type API struct {
	store *store.Store
	queue *queue.Queue[*model.Task]
}

func NewApi(s *store.Store, q *queue.Queue[*model.Task]) *API {
	return &API{
		store: s,
		queue: q,
	}
}

func (a *API) TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		a.AddTask(w, r)
	case http.MethodGet:
		a.GetTask(w, r)
	}

}
func (a *API) GetStats(w http.ResponseWriter, r *http.Request) {
	submited, completed, inProgress := a.store.GetStats()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"submited": submited, "completed": completed, "in_progress": inProgress})
}

func (a *API) AddTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		return
	}

	defer r.Body.Close()
	var data model.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	payloadData := data.Payload

	task := a.store.AddTask(payloadData)
	a.queue.Enqueue(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{

		"id": task.ID,
	})

}

func (a *API) GetTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		return
	}
	defer r.Body.Close()
	returnArray := a.store.GetTask()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnArray)
}

func (a *API) GetTaskById(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	task := a.store.GetTaskById(id)
	if task == nil {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
