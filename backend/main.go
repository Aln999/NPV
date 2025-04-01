package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/cluster-connectivity-visualizer/pkg/analyzer"
	"github.com/cluster-connectivity-visualizer/pkg/checker"
	"github.com/cluster-connectivity-visualizer/pkg/recommender"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/check", connectivityCheckHandler)
	http.HandleFunc("/api/analyze", policyAnalysisHandler)
	http.HandleFunc("/", serveFrontend)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func connectivityCheckHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	destination := r.URL.Query().Get("destination")

	result, err := checker.CheckConnectivity(clientset, source, destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recommendations := recommender.SuggestNetworkPolicy(source, destination)

	response := map[string]interface{}{
		"connectivity":    result,
		"recommendations": recommendations,
	}

	json.NewEncoder(w).Encode(response)
}

func policyAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	namespace := r.URL.Query().Get("namespace")

	analysis, err := analyzer.AnalyzePolicies(clientset, namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(analysis)
}

func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}
