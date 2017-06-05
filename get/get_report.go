package get

import (
	"agent/deal"
	//"agent/util"
	pb "agent/report"
	pm "agent/report/mess"
	seelog "github.com/cihub/seelog"
	"github.com/elastic/gosigar"
	cpu1 "github.com/jpoirier/cpu"
	"strings"
	"unicode"
)

type Host struct {
	Cpu  pm.Cpu
	Mem  pm.Mem
	Swap pm.Swap
	Disk []*pm.Disk
}

type Process struct {
	Pid       int64
	ProcMem   pm.ProcMem
	ProcState pm.ProcState
	ProcTime  pm.ProcTime
	ProcArg   pm.ProcArg
}

type MCSD struct {
	GUID       string
	ReportTime string
	Mem        pm.Mem
	Cpu        pm.Cpu
	Swap       pm.Swap
	Disk       []*pm.Disk
}

//得到主机内存的相关信息
func (self *Host) GetMem() error {
	mem := gosigar.Mem{}
	if err := mem.Get(); err != nil {
		return err
	}

	self.Mem.Total = mem.Total
	self.Mem.Used = mem.Used
	self.Mem.Free = mem.Free
	self.Mem.ActualFree = mem.ActualFree
	self.Mem.ActualUsed = mem.ActualUsed

	return nil
}

//得到cpu相关信息
func (self *Host) GetCpu() error {
	cpu := gosigar.Cpu{}
	if err := cpu.Get(); err != nil {
		return err
	}

	self.Cpu.User = cpu.User
	self.Cpu.Nice = cpu.Nice
	self.Cpu.Sys = cpu.Sys
	self.Cpu.Idle = cpu.Idle
	self.Cpu.Wait = cpu.Wait
	self.Cpu.Irq = cpu.Irq
	self.Cpu.SoftIrq = cpu.SoftIrq
	self.Cpu.Stolen = cpu.Stolen
	self.Cpu.Vender = cpu1.Vendor
	var model = strings.TrimSpace(cpu1.ProcessorFamily)
	models := strings.SplitAfter(model, "Hz")
	self.Cpu.Model = models[0]
	self.Cpu.PhysicalNumber = cpu1.Processors
	self.Cpu.LogicNumber = cpu1.OnlnProcs
	var msg = cpu1.ProcessorFamily
	fre := strings.Split(msg, "@")
	fre[1] = strings.TrimSpace(fre[1])
	fre1 := strings.SplitAfter(fre[1], "Hz")
	self.Cpu.Frequency = fre1[0]
	return nil
}

//主机虚拟内存使用情况
func (self *Host) GetSwap() error {
	swap := gosigar.Swap{}
	if err := swap.Get(); err != nil {
		return err
	}
	self.Swap.Total = swap.Total
	self.Swap.Used = swap.Used
	self.Swap.Free = swap.Free
	return nil
}

//cpu列表
func GetCpuList() error {
	cpuList := gosigar.CpuList{}
	if err := cpuList.Get(); err != nil {
		return err
	}
	return nil
}

//得到主机磁盘信息
func (self *Host) GetFileSystemUsage() {

	var diskList []*pm.Disk
	fslist := gosigar.FileSystemList{}
	fslist.Get()
	for _, fs := range fslist.List {
		var disk pm.Disk
		dir_name := fs.DirName

		usage := gosigar.FileSystemUsage{}
		usage.Get(dir_name)

		disk.DirName = dir_name
		disk.Total = usage.Total
		disk.Used = usage.Used
		disk.Avail = usage.Avail

		diskList = append(diskList, &disk)
	}
	self.Disk = diskList

}

//根据进程的pid得到进程关于状态的信息
func (self *Process) GetProcState(pid int) error {
	state := gosigar.ProcState{}
	if err := state.Get(pid); err != nil {
		return err
	}
	self.Pid = int64(pid)
	self.ProcState.Name = state.Name
	self.ProcState.Username = state.Username
	self.ProcState.Ppid = int64(state.Ppid)
	self.ProcState.Pgid = int64(state.Pgid)
	self.ProcState.Tty = int64(state.Tty)
	self.ProcState.Priority = int64(state.Priority)
	self.ProcState.Nice = int64(state.Nice)
	self.ProcState.Processor = int64(state.Processor)
	return nil
}

//根据pid得到进程的内存使用情况的相关信息
func (self *Process) GetProcMem(pid int) error {
	mem := gosigar.ProcMem{}
	if err := mem.Get(pid); err != nil {
		return err
	}
	self.ProcMem.Size = mem.Size
	self.ProcMem.Resident = mem.Resident
	self.ProcMem.Share = mem.Share
	self.ProcMem.MinorFaults = mem.MinorFaults
	self.ProcMem.MajorFaults = mem.MajorFaults
	self.ProcMem.PageFaults = mem.PageFaults
	return nil
}

func (self *Process) GetProcTime(pid int) error {

	time1 := gosigar.ProcTime{}
	if err := time1.Get(pid); err != nil {
		return err
	}
	self.ProcTime.StartTime = time1.StartTime
	self.ProcTime.User = time1.User
	self.ProcTime.Sys = time1.Sys
	self.ProcTime.Total = time1.Total
	return nil
}

func (self *Process) GetProcArgs(pid int) error {
	procArgs := gosigar.ProcArgs{}
	if err := procArgs.Get(pid); err != nil {
		return err
	}
	self.ProcArg.StartDir = strings.Join(procArgs.List, "")
	return nil
}

//方法未实现
/*func GetProcExe(pid int) error {
	procExe := gosigar.ProcExe{}
	if err := procExe.Get(pid); err != nil {
		return err
	}
	x, _ := json.Marshal(procExe)
	fmt.Println(string(x))
	return nil
}*/

func GetAgent() []string {
	agentMess := deal.ReadFile()
	return agentMess
}

//得到host信息
func GetHost() *pb.Host {
	host := Host{}
	c := ConcreteGet{}
	hostMess := pb.Host{}

	errM := host.GetMem()
	errS := host.GetSwap()
	errC := host.GetCpu()
	host.GetFileSystemUsage()
	osMess := c.GetOs()
	caption := osMess[0]
	hostName := c.GetHostName()

	var version string
	for _, v := range osMess[1] {
		if unicode.IsDigit(v) {
			version += string(v)
		}
		if strings.EqualFold(string(v), ".") {
			version += string(v)
		}
	}
	if errM != nil {
		seelog.Error("get host memory error: ", errM)
	}
	if errS != nil {
		seelog.Error("get host swap error: ", errS)
	}
	if errC != nil {
		seelog.Error("get host cpu error: ", errC)
	}
	hostMess.Swap = &host.Swap
	hostMess.Cpu = &host.Cpu
	hostMess.Mem = &host.Mem
	hostMess.Disk = host.Disk

	hostMess.Caption = caption
	hostMess.Version = version
	hostMess.HostName = hostName
	return &hostMess
}

//获取主机内存信息 5/s
func GetMCSD() *pb.MCSD {
	host := Host{}

	MCSDMess := pb.MCSD{}

	errM := host.GetMem()
	errS := host.GetSwap()
	errC := host.GetCpu()
	host.GetFileSystemUsage()
	if errM != nil {
		seelog.Error("get MCSD memory error: ", errM)
	}
	if errS != nil {
		seelog.Error("get MCSD swap error: ", errS)
	}
	if errC != nil {
		seelog.Error("get MCSD cpu error: ", errC)
	}
	MCSDMess.Swap = &host.Swap
	MCSDMess.Cpu = &host.Cpu
	MCSDMess.Mem = &host.Mem
	MCSDMess.Disk = host.Disk
	return &MCSDMess
}

func GetProByPid(pid int) pb.Process {
	process := Process{}
	processPM := pb.Process{}
	errA := process.GetProcArgs(pid)
	if errA != nil {
		//mylog.Warning("error get process args pid = ", pid)
	}
	errM := process.GetProcMem(pid)
	if errM != nil {
		//log.Println("error get process memory pid = ", pid)
	}
	errS := process.GetProcState(pid)
	if errS != nil {
		//log.Println("error get process state pid = ", pid)
	}
	errT := process.GetProcTime(pid)
	if errT != nil {
		//log.Println("error get process time pid = ", pid)
	}
	processPM.Pid = int64(pid)
	processPM.ProcArg = &process.ProcArg
	processPM.ProcMem = &process.ProcMem
	processPM.ProcState = &process.ProcState
	processPM.ProcTime = &process.ProcTime
	return processPM
}
func GetProcess() []*pb.Process {
	processMess := []*pb.Process{}
	pids := gosigar.ProcList{}
	pids.Get()

	for _, pid := range pids.List {
		process := Process{}
		processPM := pb.Process{}
		errA := process.GetProcArgs(pid)
		if errA != nil {
			//mylog.Warning("error get process args pid = ", pid)
		}
		errM := process.GetProcMem(pid)
		if errM != nil {
			//log.Println("error get process memory pid = ", pid)
		}
		errS := process.GetProcState(pid)
		if errS != nil {
			//log.Println("error get process state pid = ", pid)
		}
		errT := process.GetProcTime(pid)
		if errT != nil {
			//log.Println("error get process time pid = ", pid)
		}
		if errA != nil || errM != nil || errS != nil || errT != nil {
			continue
		}
		processPM.Pid = int64(pid)
		processPM.ProcArg = &process.ProcArg
		processPM.ProcMem = &process.ProcMem
		processPM.ProcState = &process.ProcState
		processPM.ProcTime = &process.ProcTime

		processMess = append(processMess, &processPM)
	}
	return processMess
}

func GetNetWork() *pb.NetWork {
	netWork := pb.NetWork{}
	c := ConcreteGet{}
	addrs := c.GetIp()
	portList := c.GetPort()
	netWork.Ip = addrs
	netWork.Ports = portList
	return &netWork
}
