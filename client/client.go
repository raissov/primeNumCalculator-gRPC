package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	primeNumCalculator "primeNumCalculator/proto"
	"strconv"
	"time"
)

func main() {
	log.Println("Server is starting...")
	if len(os.Args) != 2{
		log.Fatal("Number is missing!")
	}

	clientConn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to connect to the server!")
	}
	defer clientConn.Close()

	serviceClient := primeNumCalculator.NewPrimeServiceClient(clientConn)

	i, err := strconv.Atoi(os.Args[1])

	if err != nil{
		log.Fatal("Impossible to parse the number!")
	}

	request := &primeNumCalculator.PrimeRequest{
		Number: int64(i),
	}

	context, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	stream, err := serviceClient.Prime(context, request)

	if err != nil{
		log.Fatalf("Error while requestiong gRPC server: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil{
			responseErr, ok := status.FromError(err)
			if ok {
				if responseErr.Code() == codes.InvalidArgument {
					fmt.Println("Error: number must be greater than 1")
				}else if responseErr.Code() == codes.DeadlineExceeded {
					fmt.Println("Timeout")
				}else {
					log.Printf("Unknown error: %v\n", responseErr.Message())
				}
			}else {
				log.Fatalf("Error: %v", err)
			}
			break
		}
		fmt.Println(res.GetPrime())
	}
}
