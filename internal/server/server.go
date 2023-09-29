package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mxcd/go-config/config"
	"github.com/rs/zerolog/log"
)

func StartServer() {

	http.HandleFunc("/hold/", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		durationString := strings.TrimPrefix(r.URL.Path, "/hold/")
		log.Info().Msgf("Received http request to hold connection for %s seconds", durationString)

		duration, err := strconv.ParseInt(durationString, 10, 64)
		if err != nil {
			log.Warn().Err(err).Msg("Invalid duration given in request")
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Invalid duration"))
			if err != nil {
				log.Error().Err(err).Msg("Error writing response")
			}
			return
		}

		log.Info().Msgf("Holding connection for %d seconds", duration)
		time.Sleep(time.Duration(duration) * time.Second)
		log.Info().Msg("Done holding connection")
		elapsed := time.Since(startTime)

		output := fmt.Sprintf(`
			Start time: %s
			End time: %s
			Elapsed: %s
		`, startTime, time.Now(), elapsed)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))
	})

	port := config.Get().Int("PORT")
	portString := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Starting server on port %d", port)
	http.ListenAndServe(portString, nil)
}
