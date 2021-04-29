package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"ip_changer/internal/app/iptables"
	"ip_changer/internal/app/middleware"
	"log"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	config *Config
}

func newServer(config *Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		config: config,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.rootHandler()).Methods(http.MethodGet)
	s.router.HandleFunc("/healthcheck", s.healthcheckHandler()).Methods(http.MethodGet)
	s.router.Handle("/change-ip", middleware.Auth(s.changeIpHandler(), s.config.Token)).Methods(http.MethodGet)
}

func (s *server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Private property. If you are not sure - go away."))
	}
}

func (s *server) healthcheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm alive."))
	}
}

func (s *server) changeIpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, err := iptables.ChangeIp(s.config.DockerNetwork, s.config.SlackClient)
		if err != nil {
			log.Println(err)
		}
		err = json.NewEncoder(w).Encode(ip)
		if err != nil {
			log.Println(err)
		}
	}
}
