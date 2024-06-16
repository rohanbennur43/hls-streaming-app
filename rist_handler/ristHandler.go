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
	RistOutUrlChan    chan string
}

func (ristHandler *RistHandlerStruct) Init(logger *zap.SugaredLogger, RistUrlUpdateChan chan string, grpcClient grpcClient.GrpcClient, ristOutUrlChan chan string) {
	ristHandler.logger = logger
	ristHandler.RistUrlUpdateChan = RistUrlUpdateChan
	ristHandler.grpcClient = grpcClient
	ristHandler.RistOutUrlChan = ristOutUrlChan
	ristHandler.logger.Info("Successfully initialised ristHandler")

	go ristHandler.HandleRistUrlUpdate()
}

// ValidateRISTURL validates if the given URL is of type "rist://ip:port?mode=value"
func (ristHandler *RistHandlerStruct) ValidateRISTURL(ristURL string) (error, grpcClient.RistUrlStruct) {
	ristUrlstruct := grpcClient.RistUrlStruct{}
	parsedURL, err := url.Parse(ristURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err), ristUrlstruct
	}

	// Check the scheme
	if parsedURL.Scheme != "rist" {
		return fmt.Errorf("invalid scheme: %s", parsedURL.Scheme), ristUrlstruct
	}

	// Extract and validate the host (IP address)
	host, port, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return fmt.Errorf("invalid host:port format: %v", err), ristUrlstruct
	}

	// Validate the IP address
	if net.ParseIP(host) == nil {
		return fmt.Errorf("invalid IP address: %s", host), ristUrlstruct
	}

	// Validate the port number
	portPattern := `^\d+$`
	matched, err := regexp.MatchString(portPattern, port)
	if err != nil || !matched {
		return fmt.Errorf("invalid port number: %s", port), ristUrlstruct
	}

	// Extract and validate the mode query parameter
	mode := parsedURL.Query().Get("mode")
	if mode == "" {
		return fmt.Errorf("missing mode query parameter"), ristUrlstruct
	}
	ristUrlstruct.HostPort = port
	ristUrlstruct.HostUrl = host
	ristUrlstruct.Mode = mode
	ristUrlstruct.OutMode = "udp"
	ristUrlstruct.OutPort = "2222"
	ristUrlstruct.OutUrl = "127.0.0.1"
	return nil, ristUrlstruct
}

func (ristHandler *RistHandlerStruct) GetRistURLstruct(ristUrl string) (error, grpcClient.RistUrlStruct) {
	err, ristUrlStruct := ristHandler.ValidateRISTURL(ristUrl)
	if err != nil {
		return fmt.Errorf("Invalid rist url: %v", err), ristUrlStruct
	}

	return nil, ristUrlStruct
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
			err = ristHandler.grpcClient.UpdateRistUrl(ristUrlStruct)
			if err != nil {
				ristHandler.logger.Errorf("Error - %s", err)
				continue
			}
			ristOutUrl := fmt.Sprintf("%s://%s:%s", ristUrlStruct.OutMode, ristUrlStruct.OutUrl, ristUrlStruct.OutPort)
			ristHandler.RistOutUrlChan <- ristOutUrl
			// Handle the new RIST URL here
			// For example, update internal state, make gRPC calls, etc.
		default:
			ristHandler.logger.Info("Waiting for Grpc URL update")
			time.Sleep(5 * time.Second)
		}
	}
}
