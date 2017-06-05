package deal

import (
	"encoding/xml"
	seelog "github.com/cihub/seelog"
	//"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}

func CreatFile(filepath string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err1 := os.Create(filepath)
	if err1 == nil {
		seelog.Debug("create new file .i")
	} else {
		seelog.Critical("create .i failed: ", err1)
	}

	guid := GUID{}
	guidMess := guid.GetGuid()

	var result StringResources
	var rs1 ResourceString
	var rs2 ResourceString
	var rs3 ResourceString

	rs1.StringName = "GUID"
	rs1.InnerText = guidMess

	rs2.StringName = "installTime"
	rs2.InnerText = now

	rs3.StringName = "version"
	rs3.InnerText = "V1.0.1"

	//result.XMLName = ".i"
	result.ResourceString = append(result.ResourceString, rs1, rs2, rs3)
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")

	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(xml.Header)
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		//写入文件
		ioutil.WriteFile(".i", xmlOutPutData, os.ModeAppend)

		//mylog.Debug("OK~")
	} else {
		seelog.Critical("create .i failed: ", outPutErr)
	}

}
func ReadFile() []string {
	var agentMess []string
	var guid string
	var installTime string
	var version string

	content, err := ioutil.ReadFile(".i")
	if err != nil {
		seelog.Error("read file error: ", err)
		return agentMess
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		seelog.Error("read file error: ", err)
		return agentMess
	}

	for _, o := range result.ResourceString {
		switch true {
		case o.StringName == "GUID":
			guid = o.InnerText
		case o.StringName == "installTime":
			installTime = o.InnerText
		case o.StringName == "version":
			version = o.InnerText
		}
	}
	agentMess = append(agentMess, guid, installTime, version)
	return agentMess
}
func CheckFile() error {

	var FileExist = false
	//var FileName string
	osfil, _ := os.Getwd()
	osfil = filepath.ToSlash(osfil)
	dir_list, e := ioutil.ReadDir(osfil)
	if e != nil {
		seelog.Error("Run path get error")
		return e
	}
	for _, v := range dir_list {
		if strings.HasSuffix(v.Name(), ".i") {
			//FileName = v.Name()
			FileExist = true
		}
	}
	if !FileExist {

		FilePath := filepath.Join(osfil, ".i")
		CreatFile(FilePath)
	}
	return nil
}
