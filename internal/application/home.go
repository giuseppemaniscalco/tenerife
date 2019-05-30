package application

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HomeHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Home was hitted")

		w.WriteHeader(http.StatusOK)
	}
}
