package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Duration struct {
	TotalDuration string `json:"totalDuration"`
}

func createTask(duration string) {

	jsonStr := fmt.Sprintf(`{"isSync": "async", "timeDuration": "%s"}`, duration)
	jsonBody := []byte(jsonStr)

	url := "http://localhost:7777/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}

func main() {
	var wg sync.WaitGroup
	var totalDuration int64 = 0
	for i := 0; i < 50; i++ {
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGen := rand.New(randomSeed)
		randomSeconds := randomGen.Int63n(100)
		totalDuration += randomSeconds
		SECOND_TO_NANOSECOND := 1000000000
		duration := time.Duration(randomSeconds * int64(SECOND_TO_NANOSECOND))
		wg.Add(1)
		time.Sleep(250 * time.Millisecond)
		go func(duratuion time.Duration) {
			createTask(duration.String())
			wg.Done()
		}(duration)
	}
	durationFromServer, err := http.Get("http://localhost:7777/time")
	if err != nil {
		panic(err)
	}
	defer durationFromServer.Body.Close()

	data, err := io.ReadAll(durationFromServer.Body)
	if err != nil {
		panic(err)
	}

	var durationObj Duration
	jsonData := string(data)
	json.Unmarshal([]byte(jsonData), &durationObj)
	parsedDuration, err := time.ParseDuration(durationObj.TotalDuration)
	if err != nil {
		panic(err)
	}
	parsedTotalDuration := int64(parsedDuration / time.Second)
	fmt.Printf("Expected duration: %d ", totalDuration)
	fmt.Printf("Actual duration: %d", parsedTotalDuration)

	if totalDuration != parsedTotalDuration {
		notification := `Expected and actual durations dont match.
		Might be some tasks were lost
		`
		fmt.Println(notification)
	}
	wg.Wait()
}
