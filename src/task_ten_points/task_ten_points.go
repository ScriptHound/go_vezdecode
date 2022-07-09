package task_ten_points

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadTasks(filename string) string {
	durations, err := ioutil.ReadFile(filename)

	Check(err)
	return string(durations)
}

func Task(duration string) {
	startTime := time.Now()
	fmt.Printf("Task started at: %s\n", startTime)

	parsedDuration, err := time.ParseDuration(duration)
	Check(err)
	time.Sleep(parsedDuration)

	finishTime := time.Now()
	fmt.Printf("Task finished at: %s\n", finishTime)

	fmt.Printf("Total task duration: %s\n", finishTime.Sub(startTime))
}

func TenPointsMain() {
	durations := ReadTasks("tasks_durations.txt")
	splittedDurations := strings.Fields(durations)

	for _, v := range splittedDurations {
		Task(v)
	}
}
