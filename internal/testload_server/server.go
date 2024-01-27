package testload_server

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"

	"github.com/mxcd/go-config/config"
	"github.com/mxcd/tester-toolbox/internal/util"
	"github.com/rs/zerolog/log"
)

func StartServer() {

	http.HandleFunc("/load/", func(w http.ResponseWriter, r *http.Request) {
		sizeString := strings.TrimPrefix(r.URL.Path, "/load/")
		log.Info().Msgf("Received http request to send %s of data", sizeString)

		size, err := util.GetByteSizeFromString(sizeString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// w.Header().Set("Content-Type", "application/octet-stream")
		buf := make([]byte, 1024)
		var i int64 = 0
		for ; i < size; i += int64(len(buf)) {
			if size-i < int64(len(buf)) {
				buf = buf[:size-i]
			}
			_, err := rand.Read(buf)
			if err != nil {
				http.Error(w, "Error generating random data", http.StatusInternalServerError)
				return
			}
			_, err = w.Write(buf)
			if err != nil {
				return
			}
		}
	})

	port := config.Get().Int("PORT")
	portString := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Starting server on port %d", port)
	http.ListenAndServe(portString, nil)
}
