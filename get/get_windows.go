package get

import (
	"agent/util"
	"bufio"
	seelog "github.com/cihub/seelog"
	"github.com/mahonia"
	"os/exec"
	"strings"
)

func (self *Port) GetPortMsg() [][]string {

	msg := util.ExecCommand("netstat", "-ano")
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
	c := exec.Command("cmd", "ver")
	stdout, err := c.StdoutPipe()
	if err != nil {
		seelog.Error("get os message error: ", err)
	}
	c.Start()
	decoder := mahonia.NewDecoder("gb18030")
	reader := bufio.NewReader(stdout)
	line, _ := reader.ReadString('\n')
	str := decoder.ConvertString(line)

	sysmag := strings.SplitN(str, "[", -1)
	sysmag[1] = "[" + sysmag[1]
	return sysmag
}

/*func main() {
	GetPort()
}
*/
