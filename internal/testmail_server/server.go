package testmail_server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mxcd/go-config/config"
	"github.com/mxcd/tester-toolbox/internal/mail"
	"github.com/rs/zerolog/log"
)

func StartServer() {

	http.HandleFunc("/send/", func(w http.ResponseWriter, r *http.Request) {
		address := strings.TrimPrefix(r.URL.Path, "/send/")
		log.Info().Msgf("Received http request to send test mail to %s", address)
		err := mail.SendMail(address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := config.Get().Int("PORT")
	portString := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Starting server on port %d", port)
	http.ListenAndServe(portString, nil)
}
