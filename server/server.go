package server

import (
	"encoding/json"
	"igaming-service/errs"
	"igaming-service/models"
	"igaming-service/util"
	"log"
	"net/http"
	"strconv"

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

	http.HandleFunc("/payoff", server.CalculatePayoff)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) CalculatePayoff(w http.ResponseWriter, r *http.Request) {

	var body models.Configurations
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := util.GetPayoff(body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, errs.ErrInvalidConfiguration.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
