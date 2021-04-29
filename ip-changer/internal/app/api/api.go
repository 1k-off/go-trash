package api

import (
	"fmt"
	"net/http"
)

func Start(config *Config) error {
	s := newServer(config)
	s.logger.Info(fmt.Sprintf("Starting server on port %v...", config.Port))
	return http.ListenAndServe(":"+config.Port, s)
}
