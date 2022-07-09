package task_thiry_points

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	ten_points "vezdecode_go/src/task_ten_points"
)

func ThirtyPointsMain() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите максимальное количество процессоров: ")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")

	wg := sync.WaitGroup{}
	maxGoroutines, err := strconv.Atoi(text)
	ten_points.Check(err)

	durations := ten_points.ReadTasks("tasks_durations.txt")
	splittedDurations := strings.Fields(durations)

	guard := make(chan struct{}, maxGoroutines)

	for _, duration := range splittedDurations {
		guard <- struct{}{}
		wg.Add(1)
		go func(dur string) {
			ten_points.Task(dur)
			<-guard
			wg.Done()
		}(duration)
	}
	wg.Wait()

}
