package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readTasks(filename string) string {
	durations, err := ioutil.ReadFile(filename)

	check(err)
	return string(durations)
}

func task(duration string) {
	startTime := time.Now()
	fmt.Printf("Task started at: %s\n", startTime)

	parsedDuration, err := time.ParseDuration(duration)
	check(err)
	time.Sleep(parsedDuration)

	finishTime := time.Now()
	fmt.Printf("Task finished at: %s\n", finishTime)

	fmt.Printf("Total task duration: %s\n", finishTime.Sub(startTime))
}

func main() {
	durations := readTasks("tasks_durations.txt")
	splittedDurations := strings.Fields(durations)

	for _, v := range splittedDurations {
		task(v)
	}
}
