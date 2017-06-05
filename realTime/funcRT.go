package realTime

import "C"
import (
	"agent/get"
	"agent/myclinet"
	"agent/report"
	"fmt"
	seelog "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"time"
)

var c = &myclinet.Clinet

//export Fork
func Fork(pid int, ppid int) {
	now := time.Now().Format("2006-01-02 15:04:05")
	process := get.GetProByPid(pid)
	seelog.Debug("fork: ", process)
	p := report.ProcessRT{StateRT: "fork", TimeRT: now, Process: &process}
	clinet := *c
	clinet.SendReportRT(context.Background(), &p)
	fmt.Println("fork: ", p)
}

//export Exec
func Exec(tid int, pid int) {
	now := time.Now().Format("2006-01-02 15:04:05")
	process := get.GetProByPid(pid)
	seelog.Debug("exec ", process)
	p := report.ProcessRT{StateRT: "exec", TimeRT: now, Process: &process}
	clinet := *c
	clinet.SendReportRT(context.Background(), &p)
	fmt.Println("exec: ", p)
}

//export Exit
func Exit(pid int, ppid int) {
	seelog.Debug("I am a Exit pid is ", pid, " ppid is ", ppid)
	//p := get.GetProByPid(pid)
	//fmt.Println(p)
}
