package grpcClient

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RistUrlStruct struct {
	HostUrl  string
	HostPort string
}
type GrpcClient struct {
	logger     *zap.SugaredLogger
	conn       *grpc.ClientConn
	ristClient RistAppClient
}

func (grpcClient *GrpcClient) Init(logger *zap.SugaredLogger) error {
	grpcClient.logger = logger
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		grpcClient.logger.Error("Failed to create grpc conn. Returning")
		return errors.New("Failed to create grpc conn")
	}
	grpcClient.conn = conn
	grpcClient.ristClient = NewRistAppClient(conn)
	grpcClient.logger.Info("Successfully initialised grpcClient")
	return nil
}

func (grpcClient *GrpcClient) UpdateRistUrl(ristUrlStruct RistUrlStruct) error {
	grpcClient.logger.Infof("Recieved Grpc request to update rist url- %v", ristUrlStruct)
	ristUpdateconfig := RistAppconfig{InputType: "rist", InputUrl: ristUrlStruct.HostUrl, InputPort: ristUrlStruct.HostPort, OutputType: "udp", OutputPort: "9999", OutputUrl: "127.0.0.1"}
	ristStartAppResponse, err := grpcClient.ristClient.StartRistApp(context.TODO(), &ristUpdateconfig)
	if err != nil {
		grpcClient.logger.Errorf("Failed to start rist app. Err - %v", err)
		return err
	}
	grpcClient.logger.Infof("Received rist app start response - %v", ristStartAppResponse)
	return nil
}
