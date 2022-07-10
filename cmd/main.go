package main

import (
	"fmt"
	fifty_points "vezdecode_go/src/task_fifty_points"
	fourtyPoints "vezdecode_go/src/task_fourty_points"
	"vezdecode_go/src/task_ten_points"
	thirtyPoints "vezdecode_go/src/task_thirty_points"
	"vezdecode_go/src/task_twenty_points"
)

func main() {
	fmt.Println("Hello")
	task_ten_points.TenPointsMain()
	task_twenty_points.TwentyPointsMain()
	thirtyPoints.ThirtyPointsMain()
	fourtyPoints.FourtyPointsMain()
	fifty_points.FiftyPointsMain()
}
