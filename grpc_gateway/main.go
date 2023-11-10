package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	pb "github.com/techvify-oliver/protoserver/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcGW "grpc_server_testing/grpc"
	"grpc_server_testing/rabbit"
	"log"
	"net"
)

func main() {
	Server_Addr := "0.0.0.0:50080"

	go func() {
		startServer(Server_Addr)
	}()

	select {}
}

func startServer(addr string) {
	// Server
	s := grpc.NewServer()

	// Client
	Client_Addr1 := "0.0.0.0:50082"
	Client_Addr2 := "0.0.0.0:50083"
	Client_Addr3 := "0.0.0.0:50084"
	g := initGrpcClient(Client_Addr1, Client_Addr2, Client_Addr3)
	ch := InitConnectRBM()
	pb.RegisterServiceServer(s, &grpcGW.Server{
		Client: g,
		Rabbitmq: &rabbit.RabbitMQClient{
			Conn: ch,
		},
	})

	lis1, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	if err := s.Serve(lis1); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
	log.Printf("GRPC SERVER connect to %s\n", addr)
}
func initGrpcClient(addr1 string, addr2 string, addr3 string) *grpcGW.GRPCClientInit {
	var err error
	grpcConnection1, err := grpc.Dial(addr1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}
	grpcConnection2, err := grpc.Dial(addr2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	grpcConnection3, err := grpc.Dial(addr3, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	log.Printf("GRPC CLIENT connect to %s\n", addr1)
	log.Printf("GRPC CLIENT connect to %s\n", addr2)
	log.Printf("GRPC CLIENT connect to %s\n", addr3)

	return &grpcGW.GRPCClientInit{
		ServiceUser:     pb.NewServiceClient(grpcConnection1),
		ServiceContract: pb.NewServiceClient(grpcConnection2),
		ServiceRequest:  pb.NewServiceClient(grpcConnection3),
	}
}

func InitConnectRBM() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5902/")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return ch
}
