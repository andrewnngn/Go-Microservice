package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	pb "github.com/techvify-oliver/protoserver/generated"
	"go.uber.org/zap"
	"golang-boilerplate/cmd/api"
	"golang-boilerplate/config"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/logger"
	viperConfig "golang-boilerplate/internal/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

// run server with CLI
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server CLI",
	Long:  "run server with CLI",
}

// init initializes the env and logger.
func init() {
	initEnv()
	configs := viperConfig.InitConfig()
	logger := initLogger()
	grpcConn := initGrpcClient(configs.Port.GrpcClientPort)

	db := initDatabase(configs)

	//go func() {
	//	grpcServer.StartServer(db)
	//}()

	apiCmd := api.NewServerCmd(configs, logger, grpcConn, db)
	rootCmd.AddCommand(apiCmd)
}

// main is the entry point for the run command.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("run command has failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initEnv loads the. env file
func initEnv() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("skip load .env file")
		return
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("init env has failed failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

func initGrpcClient(addr string) *grpcClient.GRPCClientInit {
	var err error
	grpcConnection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	log.Printf("GRPC CLIENT connect to %s\n", addr)
	return &grpcClient.GRPCClientInit{
		Service: pb.NewServiceClient(grpcConnection),
	}
}

func initDatabase(configs *config.Configurations) *ent.Client {
	db, err := ent.Open("postgres", configs.Postgres.ConnectionString)
	if err != nil {
		fmt.Printf("init database has failed failed with error: %v\n", err)
	}
	return db
}
