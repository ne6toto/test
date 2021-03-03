package main

import ("
	"fmt"
	"github.com/shirou/gopsutil/v3/host"	
)
func main() {
	value, _ := host.SensorsTemperatures()
	fmt.Println(value)
}
