package ristHandler

import (
	"time"

	"go.uber.org/zap"
)

type RistHandlerStruct struct {
	RistUrlUpdateChan chan string
	logger            *zap.SugaredLogger
}

func (ristHandler *RistHandlerStruct) Init(logger *zap.SugaredLogger, RistUrlUpdateChan chan string) {
	ristHandler.logger = logger
	ristHandler.RistUrlUpdateChan = RistUrlUpdateChan
	ristHandler.logger.Infof("Chan is3 %+v", ristHandler.RistUrlUpdateChan)

	go ristHandler.HandleRistUrlUpdate()
}

func (ristHandler *RistHandlerStruct) HandleRistUrlUpdate() {
	ristHandler.logger.Infof("Chan is1 %+v", ristHandler.RistUrlUpdateChan)
	for {
		select {
		case newRistUrl := <-ristHandler.RistUrlUpdateChan:
			ristHandler.logger.Infof("Received new RIST URL: %s", newRistUrl)
			// Handle the new RIST URL here
			// For example, update internal state, make gRPC calls, etc.
		default:
			ristHandler.logger.Info("Waiting for Grpc URL update")
			time.Sleep(5 * time.Second)
		}
	}
}
