package store

import (
	"strconv"
	"sync"

	"github.com/DenisLi/Assignment2/internal/model"
)

type Store struct {
	mu         sync.RWMutex
	tasks      map[string]*model.Task
	nextID     int
	submitted  int
	inProgress int
	completed  int
}

func NewStore() *Store {
	return &Store{tasks: make(map[string]*model.Task)}
}

func (s *Store) createId() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	return strconv.Itoa(s.nextID)
}

func (s *Store) AddTask(payload string) *model.Task {
	s.mu.Lock()
	s.submitted++
	s.nextID++

	task := &model.Task{
		ID:      strconv.Itoa(s.nextID),
		Payload: payload,
		Status:  "PENDING",
	}
	s.tasks[task.ID] = task
	s.mu.Unlock()
	return task
}
func (s *Store) GetTask() []*model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	returnArray := make([]*model.Task, 0, len(s.tasks))
	for _, value := range s.tasks {
		returnArray = append(returnArray, value)
	}
	return returnArray
}

func (s *Store) GetTaskById(id string) *model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, ok := s.tasks[id]
	if !ok {
		return nil
	}
	return task
}

func (s *Store) GetStats() (int, int, int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.submitted, s.completed, s.inProgress
}

func (s *Store) ChangeStatus(task *model.Task, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.Status = status
	s.tasks[task.ID] = task
	switch status {
	case "IN_PROGRESS":
		s.inProgress++
	case "COMPLETED":
		s.inProgress--
		s.completed++
	}
}
