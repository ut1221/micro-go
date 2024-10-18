package service

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	v1 "micro/api/system/v1"
	"net"
	"strconv"
	"strings"
	"time"
)

type MonitorService struct {
	v1.UnimplementedMonitorServer
}

func NewMonitorService() *MonitorService {
	return &MonitorService{}
}

func (s *MonitorService) MonitorServer(ctx context.Context, req *v1.MonitorServerReq) (*v1.MonitorServerReply, error) {
	// CPU信息
	physicalCnt, _ := cpu.Counts(false)
	cpuUsage, _ := cpu.Percent(time.Second, false)
	fmt.Printf("用户CPU使用率: %.2f%%\n", cpuUsage)
	// 获取每个CPU核心的使用情况
	cpuTimes, _ := cpu.Times(false)
	// 计算用户和系统CPU使用率
	userCPUUsage := (cpuTimes[0].User / cpuTimes[0].Total()) * 100
	systemCPUUsage := (cpuTimes[0].System / cpuTimes[0].Total()) * 100

	// 内存信息
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err： ", err)
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent

	//hostInfo, _ := host.Info()
	// 获取服务器IP
	ips, err := getServerIPs()
	if err != nil {
		fmt.Println("无法获取服务器IP:", err)
	} else {
		fmt.Println("服务器IP:", ips[0])
	}
	//var sysFiles []*v1.SysFiles
	//// 获取所有分区的磁盘使用情况
	//partitions, _ := disk.Partitions(false)
	//// 遍历每个分区，获取其文件系统状态
	//for _, partition := range partitions {
	//	// 获取文件系统的使用情况
	//	usage, err := disk.Usage(partition.Mountpoint)
	//	if err != nil {
	//		fmt.Printf("无法获取分区 %s 的文件系统使用情况: %v\n", partition.Mountpoint, err)
	//		continue
	//	}
	//	var sysFile = &v1.SysFiles{
	//		DirName:     partition.Mountpoint,
	//		SysTypeName: partition.Fstype,
	//		Total:       strconv.FormatUint(usage.Total/1024/1024/1024, 10),
	//		Used:        strconv.FormatUint(usage.Used/1024/1024/1024, 10),
	//		Free:        strconv.FormatUint(usage.Free/1024/1024/1024, 10),
	//		Usage:       strconv.FormatFloat(usage.UsedPercent, 'f', 2, 64),
	//	}
	//	sysFiles = append(sysFiles, sysFile)
	//	// 打印分区信息和文件系统状态信息
	//	fmt.Printf("分区: %s\n", partition.Device)
	//	fmt.Printf("挂载点: %s\n", partition.Mountpoint)
	//	fmt.Printf("文件系统类型: %s\n", partition.Fstype)
	//	fmt.Printf("总容量: %d GB\n", usage.Total/1024/1024/1024)
	//	fmt.Printf("已用空间: %d GB\n", usage.Used/1024/1024/1024)
	//	fmt.Printf("可用空间: %d GB\n", usage.Free/1024/1024/1024)
	//	fmt.Printf("使用率: %.2f%%\n", usage.UsedPercent)
	//	fmt.Println("-------------------------------")
	//}
	return &v1.MonitorServerReply{
		Cpu: &v1.Cpu{
			CpuNum: int32(physicalCnt),
			Used:   strconv.FormatFloat(userCPUUsage, 'f', 2, 64),
			Sys:    strconv.FormatFloat(systemCPUUsage, 'f', 2, 64),
			Free:   strconv.FormatFloat(100-cpuUsage[0], 'f', 2, 64),
		},
		Mem: &v1.Mem{
			Total: strconv.FormatUint(memTotal, 10),
			Used:  strconv.FormatUint(memUsed, 10),
			Free:  strconv.FormatUint(memAva, 10),
			Usage: strconv.FormatFloat(memUsedPercent, 'f', 2, 64),
		},
		Sys: &v1.Sys{
			//OsName:       hostInfo.OS,
			//OsArch:       hostInfo.KernelArch,
			//ComputerName: hostInfo.Hostname,
			ComputerIp: strings.Join(ips, ", "),
		},
		//SysFiles: sysFiles,
	}, nil
}

// 获取服务器IP
func getServerIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return ips, nil
}
