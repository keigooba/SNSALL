package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/version", versions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defauting to port %s", port)
	}

	log.Printf("Linstening on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello cloud run")
}

func versions(w http.ResponseWriter, _ *http.Request) {
	s, err := exec.Command("gcloud", "version").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(s))
}
