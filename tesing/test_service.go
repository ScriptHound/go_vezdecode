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
	totalDuration := 0
	for i := 0; i < 100; i++ {
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGen := rand.New(randomSeed)
		randomSeconds := randomGen.Intn(600)
		totalDuration += randomSeconds
		duration := time.Duration(randomSeconds * int(time.Second))
		wg.Add(1)
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
	fmt.Println(durationObj.TotalDuration)

	wg.Wait()
}
