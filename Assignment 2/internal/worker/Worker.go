package worker

import (
	"time"

	"github.com/DenisLi/Assignment2/internal/model"
	"github.com/DenisLi/Assignment2/internal/queue"
	"github.com/DenisLi/Assignment2/internal/store"
)

func Worker(id int, Tasks *store.Store, taskQueue *queue.Queue[*model.Task], stop <-chan struct{}) {
	for {
		select {

		case <-stop:
			return
		case task, ok := <-taskQueue.Dequeue():
			if !ok {
				return
			}
			Tasks.ChangeStatus(task, "IN_PROGRESS")
			time.Sleep(2 * time.Second)
			Tasks.ChangeStatus(task, "COMPLETED")

		}
	}

}
