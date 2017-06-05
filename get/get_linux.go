package get

import (
	"agent/util"
	"bufio"
	seelog "github.com/cihub/seelog"
	"os/exec"
	"strings"
)

func (self *Port) GetPortMsg() [][]string {

	msg := util.ExecCommand("netstat", "-atunlp")
	var data [][]string
	for _, line := range msg {
		if strings.Contains(line, "LISTEN") {
			str := self.DelSpaces(line, " ")
			data = append(data, str)
		}
	}
	return data
}

func (self *OsMessage) Get() []string {
	c := exec.Command("cat", "/etc/redhat-release")
	stdout, err := c.StdoutPipe()
	if err != nil {
		seelog.Error("get os message error: ", err)
	}
	c.Start()
	reader := bufio.NewReader(stdout)
	line, _ := reader.ReadString('\n')
	sysmag := strings.SplitN(line, " ", 2)
	return sysmag
}
