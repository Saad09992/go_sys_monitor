package main

import (
	"fmt"
	"system_monitor/internal/monitor"
)

func main(){
	cpuInfo,err:=monitor.GetCpuInfo()
	if err!=nil{
		fmt.Println("Error Getting cpu info", err)
	}
	hostInfo,err:=monitor.GetHostInfo()
	if err!=nil{
		fmt.Println("Error Getting host info", err)
	}
	fmt.Printf("CPU info: %v \n", cpuInfo)
	fmt.Printf("HOST info: %v \n", hostInfo)

}