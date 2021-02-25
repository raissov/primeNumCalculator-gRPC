package main

import (
	_ "container/list"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	primeNumCalculator "primeNumCalculator/proto"
)

type server struct {}

func (*server) Prime(req *primeNumCalculator.PrimeRequest, stream primeNumCalculator.PrimeService_PrimeServer)error{
	log.Printf("Received Prime gRPC request with number to decompose: %d\n", req.GetNumber())

	number := req.GetNumber()

	if number < 2{
		return status.Errorf(codes.InvalidArgument, "Error: number is less than 1, must be more")
	}

	x := int64(2)

	for number > 1 {
		if number%x == 0 {
			res := &primeNumCalculator.PrimeResponse{
				Prime: x,
			}
			err := stream.Send(res)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			number = number / x
		}else {
			x = x + 1
		}
	}
	return nil
}

func main(){
	log.Println("Starting gRPC server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Unable to connecto to the port: %v", err)
	}

	defer lis.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	primeNumCalculator.RegisterPrimeServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}