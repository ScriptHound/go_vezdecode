# Go async task queue

This is a project for Go task from VK hackathon 

Got 120 points of 150 on hackathon.

10 points task is completed in src/task_ten_points/task_ten_points.go

20 points task is completed in src/task_twenty_points/task_twenty_points.go

30 points task is completed in src/task_thirty_points/task_thirty_points.go

40 points task is completed in src/task_fourty_points/task_fourty_points.go

50 points task is completed in src/task_fifty_points/task_fifty_points.go

# Deployment

```bash
sudo apt-get install build-essential
make set_env
go mod tidy
```

In case of GO111MODULE error just
```bash
export GO111MODULE=on
```

# Run service

```bash
make run
```

# Build service
```bash
make build
```

To launch each task separately comment an appropriate line in cmd/main.go

# Testing 40 points task

```bash
go run testing/test_service.go
```
