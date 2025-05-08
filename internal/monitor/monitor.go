package monitor

import (
	"fmt"
	"time"

	"github.com/DataDog/gopsutil/cpu"
	"github.com/DataDog/gopsutil/host"
)

type HostInfo struct{
	host string
	uptime int64
	os string
	id string
}

type CpuInfo struct{
	core int
	model string
	usage int64
}

// strconv.FormatUint(vmStat.Free/megabyteDiv, 10)

const megabyteDiv uint64 = 1024 * 1024
const gigabyteDiv uint64 = megabyteDiv * 1024

func GetHostInfo()(HostInfo, error){
	var hostInfo HostInfo
	info,err:=host.Info()
	if err!=nil{
		fmt.Printf("Error fetching host info, %v", err)
		return hostInfo,err
	}
	hostInfo.host=info.Hostname
	hostInfo.uptime=int64(info.Uptime)
	hostInfo.os=info.Platform
	hostInfo.id=info.HostID
	return hostInfo, nil
}

func GetCpuInfo()([]CpuInfo, error){
	var cpuInfo []CpuInfo
	info,err:=cpu.Info()
	if err!=nil{
		fmt.Printf("Error fetching cpu info, %v", err)
		return cpuInfo,nil
	}
	usage,err:=cpu.Percent(time.Second,true)
	if err!=nil{
		fmt.Printf("Error fetching cpu usage, %v \n", err)
		return cpuInfo,nil
	}


	for i,val:=range info{
		info:=CpuInfo{
			core: int(val.CPU),
			model: val.ModelName,
			usage: int64(usage[i]),

		}
		fmt.Printf("CPU info: %v \n", info)
		cpuInfo=append(cpuInfo, info)
	}

	
	return cpuInfo,nil
}