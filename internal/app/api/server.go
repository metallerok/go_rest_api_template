package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	cfg    *ini.File
}

func NewServer(cfg *ini.File) *server {

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log_level, err := logrus.ParseLevel(cfg.Section("server").Key("log_level").String())
	if err != nil {
		log_level = logrus.DebugLevel
	}
	logger.SetLevel(log_level)

	serv := &server{
		router: mux.NewRouter(),
		logger: logger,
		cfg:    cfg,
	}

	serv.configureRouter()

	return serv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Logger() *logrus.Logger {
	return s.logger
}

func (s *server) Start() error {
	return http.ListenAndServe(s.cfg.Section("server").Key("listen_addr").String(), s)
}

func (s *server) configureRouter() {
	s.router.Use(AuthMiddleware)

	s.router.HandleFunc("/", handleDefault()).Methods("GET")
}
