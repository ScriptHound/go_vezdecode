package task_fourty_points

import (
	"encoding/json"
	"sync"
	"time"
	ten_points "vezdecode_go/src/task_ten_points"
)

type SafeQueue struct {
	mu    sync.Mutex
	queue []Task
}

func (q *SafeQueue) Append(task Task) {
	q.mu.Lock()
	q.queue = append(q.queue, task)
	q.mu.Unlock()
}

func (q *SafeQueue) Pop() Task {
	q.mu.Lock()
	task := pop(&q.queue)
	q.mu.Unlock()
	return task
}

func (q *SafeQueue) Len() int {
	return len(q.queue)
}

func (q *SafeQueue) ListAllTasks() string {
	tasks := q.queue
	serialized, err := json.Marshal(tasks)
	ten_points.Check(err)
	return string(serialized)
}

func (q *SafeQueue) GetTotalTaskDuration() string {
	var sum int64 = 0
	for _, task := range q.queue {
		duration, err := time.ParseDuration(task.Duration)
		ten_points.Check(err)
		sum += int64(duration)
	}
	duration := time.Duration(sum)
	return duration.String()
}

func pop(a *[]Task) Task {
	result := (*a)[0]
	*a = (*a)[1:]
	return result
}
