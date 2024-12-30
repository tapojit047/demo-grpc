package main

import (
	"context"
	"github.com/tapojit047/demo-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
		return
	}
	serverRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
		return
	}
}
