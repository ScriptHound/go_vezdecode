package task_fourty_points

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ten_points "vezdecode_go/src/task_ten_points"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Task struct {
	UUID     string `json:"UUID"`
	Duration string `json:"timeDuration"`
}

type DecodedTask struct {
	IsSync       string `json:"isSync"`
	TimeDuration string `json:"timeDuration"`
}

var taskChan = make(chan Task)

var queue = SafeQueue{queue: make([]Task, 0)}

func worker() {
	for {
		if queue.Len() > 0 {
			task := queue.Pop()
			duration := task.Duration
			ten_points.Task(duration)
		}
	}
}

func queueHandler(taskChan chan Task) {
	for {
		select {
		case task, ok := <-taskChan:
			if ok {
				queue.Append(task)
			} else {
				fmt.Println("Channel closed!")
			}
		default:
		}
	}
}

func nonFreeze(w http.ResponseWriter, r *http.Request) {
	var decodedTask DecodedTask
	bodyBytes, err := ioutil.ReadAll(r.Body)
	ten_points.Check(err)

	err = json.Unmarshal([]byte(bodyBytes), &decodedTask)
	ten_points.Check(err)

	if decodedTask.IsSync == "async" {
		uuid := uuid.New().String()
		duration := decodedTask.TimeDuration
		task := Task{UUID: uuid, Duration: duration}
		taskChan <- task
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(uuid))
		ten_points.Check(err)
	} else if decodedTask.IsSync == "sync" {
		for {
			if queue.Len() == 0 {
				break
			}
		}
		ten_points.Task(decodedTask.TimeDuration)
		w.WriteHeader(http.StatusOK)
	}
}

func getScheduledTasks(w http.ResponseWriter, r *http.Request) {
	scheduledTasks := queue.ListAllTasks()
	_, err := w.Write([]byte(scheduledTasks))
	ten_points.Check(err)
	w.WriteHeader(http.StatusOK)
}

func getTotalTasksDuration(w http.ResponseWriter, r *http.Request) {
	totalTime := queue.GetTotalTaskDuration()
	totalTime = fmt.Sprintf(`{"totalDuration": "%s"}`, totalTime)
	_, err := w.Write([]byte(totalTime))
	ten_points.Check(err)
	w.WriteHeader(http.StatusOK)
}

func FourtyPointsMain() {
	go queueHandler(taskChan)
	go worker()
	router := mux.NewRouter()
	router.HandleFunc("/add", nonFreeze).Methods("POST")
	router.HandleFunc("/schedule", getScheduledTasks).Methods("GET")
	router.HandleFunc("/time", getTotalTasksDuration).Methods("GET")
	http.Handle("/", router)
	err := http.ListenAndServe(":7777", nil)
	ten_points.Check(err)
}
