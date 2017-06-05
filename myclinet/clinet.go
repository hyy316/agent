package myclinet

import (
	pb "agent/report"
	seelog "github.com/cihub/seelog"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn
var Clinet pb.GreeterClient
var err error

func GetConn(address string) {
	//flag.Parse()
	//address = *Ip + ":" + *port
	seelog.Debug("connect to ", address, "......")
	Conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		seelog.Error("connect failed: ", err)
	}

}

func Get(address string) {
	GetConn(address)
	Clinet = pb.NewGreeterClient(Conn)
	//RealTime()
}
