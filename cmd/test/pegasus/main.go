package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"time"

	"github.com/XiaoMi/pegasus-go-client/pegasus"
	pb "github.com/mloves0824/antalk_web/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) TestSet(ctx context.Context, in *pb.SetReq) (*pb.SetReply, error) {

	cfgPath, _ := filepath.Abs("./config/pegasus-client-config.json")
	rawCfg, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		//fmt.Println(err)
		return &pb.SetReply{Status: -1}, nil
	}

	cfg := &pegasus.Config{}
	json.Unmarshal(rawCfg, cfg)
	c := pegasus.NewClient(*cfg)

	tb, err := c.OpenTable(context.Background(), "temp")
	if err != nil {
		return &pb.SetReply{Status: -2}, nil
	}

	time.Sleep(time.Second * 300)

	value := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		value[i] = 'x'
	}
	//in.Key
	tb.Set(context.Background(), []byte("hash"), []byte("sort"), value)

	return &pb.SetReply{Status: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestPegasusServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
