package ristHandler

import (
	"fmt"
	"hls-streamer/grpcClient"
	"net"
	"net/url"
	"regexp"
	"time"

	"go.uber.org/zap"
)

type RistHandlerStruct struct {
	RistUrlUpdateChan chan string
	logger            *zap.SugaredLogger
	grpcClient        grpcClient.GrpcClient
}

func (ristHandler *RistHandlerStruct) Init(logger *zap.SugaredLogger, RistUrlUpdateChan chan string, grpcClient grpcClient.GrpcClient) {
	ristHandler.logger = logger
	ristHandler.RistUrlUpdateChan = RistUrlUpdateChan
	ristHandler.grpcClient = grpcClient
	ristHandler.logger.Info("Successfully initialised ristHandler")

	go ristHandler.HandleRistUrlUpdate()
}
func (ristHandler *RistHandlerStruct) ValidateRISTURL(ristURL string) (error, string, string) {
	parsedURL, err := url.Parse(ristURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err), "", ""
	}

	// Check the scheme
	if parsedURL.Scheme != "rist" {
		return fmt.Errorf("invalid scheme: %s", parsedURL.Scheme), "", ""
	}

	// Extract and validate the host (IP address)
	host, port, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return fmt.Errorf("invalid host:port format: %v", err), "", ""
	}

	// Validate the IP address
	if net.ParseIP(host) == nil {
		return fmt.Errorf("invalid IP address: %s", host), "", ""
	}

	// Validate the port number
	portPattern := `^\d+$`
	matched, err := regexp.MatchString(portPattern, port)
	if err != nil || !matched {
		return fmt.Errorf("invalid port number: %s", port), "", ""
	}

	return nil, host, port
}

func (ristHandler *RistHandlerStruct) GetRistURLstruct(ristUrl string) (error, grpcClient.RistUrlStruct) {
	ristUrlstruct := grpcClient.RistUrlStruct{}
	err, ristHost, ristPort := ristHandler.ValidateRISTURL(ristUrl)
	if err != nil {
		return fmt.Errorf("Invalid rist url: %v", err), ristUrlstruct
	}
	ristUrlstruct.HostPort = ristPort
	ristUrlstruct.HostUrl = ristHost
	return nil, ristUrlstruct
}

func (ristHandler *RistHandlerStruct) HandleRistUrlUpdate() {
	for {
		select {
		case newRistUrl := <-ristHandler.RistUrlUpdateChan:
			ristHandler.logger.Infof("Received new RIST URL: %s", newRistUrl)
			err, ristUrlStruct := ristHandler.GetRistURLstruct(newRistUrl)
			if err != nil {
				ristHandler.logger.Errorf("Ignoring rist url %s. Err- %v", newRistUrl, err)
				continue
			}
			ristHandler.grpcClient.UpdateRistUrl(ristUrlStruct)
			// Handle the new RIST URL here
			// For example, update internal state, make gRPC calls, etc.
		default:
			ristHandler.logger.Info("Waiting for Grpc URL update")
			time.Sleep(5 * time.Second)
		}
	}
}
