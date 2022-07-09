package task_twenty_points

import (
	"fmt"
	"strings"
	"sync"
	"time"
	ten_points "vezdecode_go/src/task_ten_points"
)

func Task(duration string, wg *sync.WaitGroup) {
	startTime := time.Now()
	startString := fmt.Sprintf("Task started at: %s\n", startTime)
	fmt.Println(startString)

	parsedDuration, err := time.ParseDuration(duration)
	ten_points.Check(err)
	time.Sleep(parsedDuration)

	finishTime := time.Now()
	finishString := fmt.Sprintf("Task finished at: %s\n", finishTime)
	fmt.Println(finishString)

	durString := fmt.Sprintf("Total task duration: %s\n", finishTime.Sub(startTime))
	fmt.Println(durString)
	wg.Done()
}

func TwentyPointsMain() {
	durations := ten_points.ReadTasks("tasks_durations.txt")
	splittedDurations := strings.Fields(durations)
	var wg sync.WaitGroup
	for _, v := range splittedDurations {
		wg.Add(1)
		go Task(v, &wg)
	}
	wg.Wait()
}
