package server

import (
	"net/http"

	"go.uber.org/zap"
)

type ServerStruct struct {
	logger            *zap.SugaredLogger
	ristUrlUpdateChan chan string
}

func (s *ServerStruct) Init(logger *zap.SugaredLogger, ristUrlUpdateChan chan string) {
	s.logger = logger
	s.ristUrlUpdateChan = ristUrlUpdateChan
}

func (s *ServerStruct) UpdateRistUrl(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Received UpdateUrl request")
	if r.Method != http.MethodPut {
		s.logger.Infof("Received request is not of type %s. Ignoring Request", http.MethodPut)
		return
	}

	ristUrl := r.URL.Query().Get("rist_url")
	s.logger.Infof("Received rist url update request. New rist url is %s. Sending the update to grpc handler", ristUrl)
	s.logger.Infof("Chan is %+v", s.ristUrlUpdateChan)
	s.ristUrlUpdateChan <- ristUrl
	return

}
