package main

import (
	pb "agent/report"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SendReport(ctx context.Context, in *pb.Report) (*pb.Reply, error) {
	//log.Println(in.Process)
	log.Println("get all msg----", in.GUID, in.ReportTime, in.Host.Caption, in.Host.Version,in.NetWork)
	
	return &pb.Reply{Mess: 1}, nil
}

func (s *server) SendMCSD(ctx context.Context, in *pb.MCSD) (*pb.ReplyMCSD, error) {
	//log.Println(in.Process)
	log.Println("get MCSD msg----", in.GUID, in.ReportTime)
	return &pb.ReplyMCSD{Mess: 100}, nil
}

func (s *server) SendReportRT(ctx context.Context, in *pb.ProcessRT) (*pb.ReplyRT, error) {
	log.Println(in)
	return &pb.ReplyRT{Mess: 200}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
