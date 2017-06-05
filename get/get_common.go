package get

import (
	pm "agent/report/mess"
	"log"
	"net"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (self *Port) PortList(data [][]string) []*pm.Port {
	var list []*pm.Port
	for _, str := range data {
		var port pm.Port
		port.Protocol = str[0]
		s := self.DelSpaces(str[1], ":")
		var str1 int64
		if len(s) == 1 {
			str1, _ = strconv.ParseInt(s[0], 10, 64)
		}
		if len(s) > 1 {
			str1, _ = strconv.ParseInt(s[1], 10, 64)
		}
		port.PortNumber = str1
		str[4] = strings.Replace(str[4], "\r", "", -1)
		str4, _ := strconv.ParseInt(str[4], 10, 64)
		port.Pid = str4
		list = append(list, &port)
	}
	return list
}
func (self *Port) DelSpaces(str string, del string) []string {
	var s = strings.Split(str, del)
	var data []string
	for _, str := range s {
		if utf8.RuneCountInString(str) > 1 {
			data = append(data, str)
		}
	}
	return data
}

func (self *Ip) Get() []string {
	ipAddrs := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		//os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddrs = append(ipAddrs, ipnet.IP.String())
			}

		}
	}
	//log.Println(ipAddrs)
	return ipAddrs

}
