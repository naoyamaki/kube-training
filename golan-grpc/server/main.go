package main

import (
	"context"
	"fmt"
	"golan-grpc/pb"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(cpx context.Context, req *pb.ListFileRequest) (*pb.ListFileResponse, error) {
	fmt.Println("listfile war invoked")

}
