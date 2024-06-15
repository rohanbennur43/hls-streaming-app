package router

import (
	"hls-streamer/server"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type RouterStruct struct {
	Router *mux.Router
	logger *zap.SugaredLogger
}

func (r *RouterStruct) Init(logger *zap.SugaredLogger) {
	r.logger = logger
	r.Router = mux.NewRouter()
}

func (r *RouterStruct) HandleRoutes(s *server.ServerStruct) {
	r.logger.Info("Handling Routes")
	r.Router.HandleFunc("/update", s.UpdateRistUrl).Methods("PUT")
}
