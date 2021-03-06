package main

import (
	"bufio"
	"flag"
	log "github.com/golang/glog"
	pb "github.com/zhuanxuhit/go-in-practice/wheel/rpc/grpc/v1/tutorial"
	"github.com/zhuanxuhit/go-in-practice/wheel/rpc/grpc/v1/util"
	"net"
)

func main() {
	flag.Parse()
	defer log.Flush()
	raddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatalln("ResolveTCPAddr error:", err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatalln("DialTCP error:", err)
	}
	var (
		rr = bufio.NewReader(conn)
		wr = bufio.NewWriter(conn)
	)
	// 1. 发送数据
	message := &pb.Message{
		Request: &pb.Request{
			Request1: &pb.XX1Request{
				Name: "XX1Request",
			},
		},
	}
	log.Info("begin to WriteTCP message")
	err = util.WriteTCP(wr, message)
	if err != nil {
		log.Fatalln("writeTCP error:", err)
	}
	// 2. 接收数据
	newMesage, err := util.ReadTCP(rr)
	if err != nil {
		log.Fatalln("readTCP error:", err)
	}
	log.Info("read message:", newMesage)
	log.Info("client exist")
}
