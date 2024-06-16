package router

import (
	"hls-streamer/server"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type RouterStruct struct {
	Router *mux.Router
	logger *zap.SugaredLogger
}

func (r *RouterStruct) Init(logger *zap.SugaredLogger) {
	r.logger = logger
	r.logger.Info("Successfully initialised router")

	r.Router = mux.NewRouter()
}

func (r *RouterStruct) HandleRoutes(s *server.ServerStruct) {
	r.logger.Info("Handling Routes")
	r.Router.HandleFunc("/update", s.UpdateRistUrl).Methods("PUT")
	r.Router.HandleFunc("/", s.FrontEndPage).Methods("GET")
	r.Router.PathPrefix("/hls-stream-files/").Handler(http.StripPrefix("/hls-stream-files/", http.FileServer(http.Dir("hls-stream-files/")))).Methods("GET")

}
