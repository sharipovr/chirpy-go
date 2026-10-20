package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println(err)
	}
}

func (cfg *apiConfig) writehitsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Hits: %v", cfg.fileserverHits.Load())))
	if err != nil {
		fmt.Println(err)
	}
}

func (cfg *apiConfig) resethitsHnadler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)
}

func main() {
	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	var apiCfg apiConfig
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(appHandler))
	mux.HandleFunc("/healthz", readinessHandler)
	mux.HandleFunc("/metrics", apiCfg.writehitsHandler)
	mux.HandleFunc("/reset", apiCfg.resethitsHnadler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
