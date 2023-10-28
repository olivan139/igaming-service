package server

import (
	"encoding/json"
	"igaming-service/errs"
	"igaming-service/models"
	"igaming-service/util"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	Port int
}

func Run(port int) {
	log.Printf("Server started on port: %v\n", port)

	server := &Server{
		Port: port,
	}
	r := prometheus.NewRegistry()
	r.MustRegister(requestDuration, badRequestsCounter, totalRequestsCounter)
	http.HandleFunc("/payoff", server.CalculatePayoff)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) CalculatePayoff(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		method := r.Method
		elapsed := time.Since(start).Seconds()
		requestDuration.WithLabelValues(method).Observe(elapsed)
		totalRequestsCounter.Inc()
	}()

	var body models.Configurations
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		badRequestsCounter.Inc()
		return
	}

	response, err := util.GetPayoff(body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, errs.ErrInvalidConfiguration.Error(), http.StatusBadRequest)
		badRequestsCounter.Inc()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		badRequestsCounter.Inc()
	}
}
