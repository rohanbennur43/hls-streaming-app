package main

import (
	"hls-streamer/grpcClient"
	"hls-streamer/logging"
	ristHandler "hls-streamer/rist_handler"
	"hls-streamer/router"
	"hls-streamer/server"
	"log"
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var logErr error

func main() {
	_, logger, logErr = logging.LogInit("streamer_app")
	if logErr != nil {
		//using default logger (log)
		log.Println("failed to initialize logger in main module, err =", logErr.Error())
		panic(logErr)
	}

	ristUrlUpdateChan := make(chan string)

	grpcClient := grpcClient.GrpcClient{}
	grpcClient.Init(logger)

	ristHandler := ristHandler.RistHandlerStruct{}
	ristHandler.Init(logger, ristUrlUpdateChan, grpcClient)

	server := server.ServerStruct{}
	server.Init(logger, ristUrlUpdateChan)

	router := router.RouterStruct{}
	router.Init(logger)

	router.HandleRoutes(&server)

	http.ListenAndServe("127.0.0.1:8888", router.Router)

}
