package webvitals_exporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/vitals", HandleWebVital)
	mux.Handle("/metrics", promhttp.Handler())
	fmt.Println("running server on " + port)
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		panic(err)
	}
}
