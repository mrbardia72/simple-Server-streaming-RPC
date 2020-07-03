package main

import (
   "context"
   "fmt"
   "../calculatorpb"
   "google.golang.org/grpc"
   "io"
   "log"
)

func main() {
   fmt.Println("Client is running...")

   cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
   if err != nil {
      log.Fatalf("could not connect to server: %v", err)
   }
   defer cc.Close()

   c := calculatorpb.NewCalculatorServiceClient(cc)

   doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
   fmt.Println("Starting to do a PrimeDecomposition server streaming RPC")

   req := &calculatorpb.PrimeNumberDecompositionRequest{
      Number: 120,
   }

   stream, err := c.PrimeNumberDecomposition(context.Background(), req)
   if err != nil {
      log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
   }

   for {
      res, err := stream.Recv()
      if err == io.EOF {
         break
      }
      if err != nil {
         log.Printf("Error while streaming PrimeDecomposition RPC: %v", err)
      }
      fmt.Println(res.PrimeFactor)
   }
}