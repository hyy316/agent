package main

import (
"fmt"
	"flag"
	"sync"
	"time"
	"github.com/kardianos/service"
	"os"
	"agent/deal"
	"agent/get"
	"agent/myclinet"
	pb "agent/report"
	seelog "github.com/cihub/seelog"
	"github.com/robfig/cron"
	"golang.org/x/net/context"
	"agent/util"
	"path/filepath"
	 "os/exec"
    "strings"
)

type ReportMess struct {
	PbReport pb.Report
	PbMCSD   pb.MCSD
}

var runState = "running"

var waitgroup sync.WaitGroup
var Ip *string = flag.String("I", "localhost", "Use -I <IP address>")
var port *string = flag.String("P", "50051", "Use -P <port>")

//var conn *grpc.ClientConn
var clinet pb.GreeterClient
var err error

func initdemo() {
  	file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    splitstring := strings.Split(path, "\\")
    size := len(splitstring)
    splitstring = strings.Split(path, splitstring[size-1])
    ret := strings.Replace(splitstring[0], "\\", "/", size-1)

    fmt.Println(ret)

	add := ret + "seelog.xml"

	logger, err := seelog.LoggerFromConfigAsFile(add)

	if err != nil {
		seelog.Critical("parsing config log file error: ", err)
		return
	}
	seelog.ReplaceLogger(logger)

}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error { 
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}


func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}



func main() {


	address,_:= os.Getwd()

	address=filepath.ToSlash(address)

	util.AlterXML(address)

	initdemo()




	//服务的配置信息
	cfg := &service.Config{
		Name:        "agent",
		DisplayName: "a agent service",
		Description: "This is a Go service.",
		//Executable:  "D:/Workspaces/go/src/agent/cyclereport.exe",
		//WorkingDirectory: "D:/mygo/src/agent",
		//ChRoot: "D:/mygo/src/agent/cyclereport.exe",
	}
	// Interface 接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		seelog.Debug(err)
	}
	// logger 用于记录系统日志
	/*logger, err := s.Logger(nil)
	if err != nil {
		seelog.Debug(err)
	}*/
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			seelog.Debug(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			seelog.Debug(err)
		}
	}
	if err != nil {
		seelog.Debug(err)
	}
}



func (p *program) run() {

	defer seelog.Flush()
	deal.CheckFile()
	flag.Parse()
	address := *Ip + ":" + *port
	myclinet.Get(address)

	setUp()
	seelog.Debug("report start------")
	report()
	time.Sleep(10 * time.Second)
	waitgroup.Add(2)
	go reportLong()
	go reportShort()
	//go reportRT()
	waitgroup.Wait()
}

func setUp() {
	clinet = myclinet.Clinet
}

func reportLong() {
	spec := "0, 0/1, *, *, ?, *"
	c := cron.New()
	c.AddFunc(spec, report)
	c.Start()
	select {}
}

func reportShort() {

	spec := "0/5, *, *, *, ?, *"
	c := cron.New()
	c.AddFunc(spec, reportMCSD)
	c.Start()
	select {}
}

/*func reportRT() {
	realTime.ReportRT()
}
*/

func report() {
	errR := send()
	if errR != nil {
		seelog.Error("send all report error:	", errR)
		retry(3, send)
	}
}

func reportMCSD() {
	errR := sendMCSD()
	if errR != nil {
		seelog.Error("send MCSD report error:	", errR)
		retry(3, send)
	}
}

//上报信息
func send() error {
	rm := ReportMess{}
	rm.getMessage()
	seelog.Debug("send all report  ......")
	r, errSR := clinet.SendReport(context.Background(), &rm.PbReport)
	if errSR != nil {
		seelog.Errorf("report could not greet: %v", errSR)
		runState = "connect failed"
		seelog.Error("state:	", "(runState: ", runState, ")")
		//mylog.Error("could not greet: %v", errSR)
		return errSR

	}
	seelog.Debug("report greeting: ", r.Mess)
	runState = "running"
	seelog.Debug("state:	", "(runState: ", runState, ")")
	//mylog.Debug("Greeting: ", r.Mess)
	return nil
}

func sendMCSD() error {
	rm := ReportMess{}
	rm.getMCSDMessage()
	seelog.Debug("send MCSD report  ......")
	r, errSR := clinet.SendMCSD(context.Background(), &rm.PbMCSD)
	if errSR != nil {
		seelog.Errorf("MCSD could not greet: %v", errSR)
		runState = "connect failed"
		seelog.Error("state:	", "(runState: ", runState, ")")
		//mylog.Error("could not greet: %v", errSR)
		return errSR

	}
	seelog.Debug("MCSD greeting: ", r.Mess)
	//mylog.Debug("Greeting: ", r.Mess)
	runState = "running"
	seelog.Debug("state:	", "(runState: ", runState, ")")
	return nil
}

//获取所需上报信息
func (rm *ReportMess) getMessage() {
	var guid string
	var agent pb.Agent

	agentMess := get.GetAgent()
	//fmt.Println(agentMess[0])
	if len(agentMess) >= 3 {
		guid = agentMess[0]
		agent.InstallTime = agentMess[1]
		agent.Version = agentMess[2]
	}

	hostMess := get.GetHost()
	processMess := get.GetProcess()
	netWork := get.GetNetWork()

	rm.PbReport.GUID = guid
	rm.PbReport.Host = hostMess
	rm.PbReport.Process = processMess
	rm.PbReport.NetWork = netWork
	rm.PbReport.Agent = &agent
	now := time.Now().Format("2006-01-02 15:04:05")
	rm.PbReport.ReportTime = now
}

func (rm *ReportMess) getMCSDMessage() {
	var guid string

	agentMess := get.GetAgent()
	//fmt.Println(agentMess[0])
	if len(agentMess) >= 3 {
		guid = agentMess[0]
		//agent.InstallTime = agentMess[1]
		//agent.Version = agentMess[2]
	}

	MCSDMess := get.GetMCSD()
	rm.PbMCSD.GUID = guid 
	rm.PbMCSD.Mem = MCSDMess.Mem
	rm.PbMCSD.Cpu = MCSDMess.Cpu
	rm.PbMCSD.Swap = MCSDMess.Swap
	rm.PbMCSD.Disk = MCSDMess.Disk
	now := time.Now().Format("2006-01-02 15:04:05")
	rm.PbMCSD.ReportTime = now
}

//重试
func retry(attempts int, callback func() error) (errR error) {
	for i := 1; ; i++ {
		errR := callback()
		if errR == nil {
			return errR
		}

		if i > attempts {
			break
		}

		time.Sleep(5 * time.Second)

		seelog.Debug("the	", i, "	time retrying...")
		seelog.Error("error: ", errR)
	}
	seelog.Error("error:after ", attempts, " attempts, last error: ", errR)

	return errR
}
