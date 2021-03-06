package main

import (
   "fmt"
   "../calculatorpb"
   "google.golang.org/grpc"
   "log"
   "net"
)

type server struct{} //routing

func main() {
   fmt.Println("Server is running...")

   // Make a listener
   lis, err := net.Listen("tcp", "0.0.0.0:50051")
   if err != nil {
      log.Fatalf("Failed to listen: %v", err)
   }

   // Make a gRPC server
   grpcServer := grpc.NewServer()
   calculatorpb.RegisterCalculatorServiceServer(grpcServer, &server{})

   // Run the gRPC server
   if err := grpcServer.Serve(lis); err != nil {
      log.Fatalf("Failed to serve: %v", err)
   }
}
 
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
   fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

   number := req.Number
   divisor := int64(2)

   for number > 1 {
      if number % divisor == 0 {
         err := stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
            PrimeFactor: divisor,
         })
         if err != nil {
            log.Fatalf("Failed to send response: %v\n", err)
         }

         number = number / divisor
      } else {
         divisor++
         fmt.Printf("Divisor has increased to %v", divisor)
      }
   }

   return nil
}