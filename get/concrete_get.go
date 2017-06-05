package get

import (
	pm "agent/report/mess"
	seelog "github.com/cihub/seelog"
	"os"
)

type ConcreteGet struct{}

func (c *ConcreteGet) GetPort() []*pm.Port {
	p := Port{}
	data := p.GetPortMsg()
	portList := p.PortList(data)
	//fmt.Print(portList)
	return portList
}
func (c *ConcreteGet) GetIp() []string {
	i := Ip{}
	ips := i.Get()
	return ips
}

func (c *ConcreteGet) GetOs() []string {
	o := OsMessage{}
	sysmag := o.Get()
	return sysmag
}
func (c *ConcreteGet) GetHostName() string {
	var hostname string
	host, err := os.Hostname()
	if err != nil {
		seelog.Error("get host name error: ", err)
	} else {
		hostname = host
	}
	return hostname
}
