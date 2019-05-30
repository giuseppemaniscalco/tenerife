package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LivenessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Liveness was hitted")

		w.WriteHeader(http.StatusOK)
	}
}

func ReadynessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Readyness was hitted")

		w.WriteHeader(http.StatusOK)
	}
}
