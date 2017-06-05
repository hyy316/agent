package util

import (
	"encoding/xml"
	"io/ioutil"
	seelog "github.com/cihub/seelog"
	"os"
)

type Seelog struct{
	XMLName xml.Name `xml:"seelog"`
	Outputs Outputs `xml:"outputs"`
	Formats []Formats `xml:"formats"`
} 

type Outputs struct{
	XMLName xml.Name `xml:"outputs"`
	Filter []Filter `xml:"filter"`
	FormatId string `xml:"formatid,attr"`
}

type Formats struct{
	XMLName xml.Name `xml:"formats"`
	Format Format `xml:"format"`
}

type Filter struct{
	XMLName xml.Name `xml:"filter"`
	File File `xml:"file"`
	Levels string `xml:"levels,attr"`
}

type File struct{
	XMLName xml.Name `xml:"file"`
	Path string `xml:"path,attr"`
}

type Format struct{
	XMLName xml.Name `xml:"format"`
	ID string `xml:"id,attr"`
	Formatdata string `xml:"format,attr"`
}

func AlterXML(address string){

	address_xml:=address+"/seelog.xml"
	address_log:=address+"/agent.log"

	content,err :=ioutil.ReadFile(address_xml)
	if err!=nil{
		seelog.Error(err)
	}

	var result Seelog

	err =xml.Unmarshal(content,&result)

	if err!=nil{
		seelog.Error(err)
	}
	for index,_:=range result.Outputs.Filter{
			result.Outputs.Filter[index].File.Path=address_log
	}

		//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		/*headerBytes := []byte(xml.Header)
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)*/
		//写入文件
		ioutil.WriteFile(address_xml, xmlOutPut, os.ModeAppend)
	} else {
		seelog.Error(outPutErr)
	}

}