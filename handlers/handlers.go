package handlers

import (
	"crocStuff/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func ClientPushing(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit ClientPushing Post Endpoint.")

	var req models.ClientPushingRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.ClientPushingResponse{Status: "success"}

	crocExecutable, _ := exec.LookPath("croc")
	cmdGoVer := &exec.Cmd{
		Path:   crocExecutable,
		Args:   []string{crocExecutable, req.CrocCodePhrase},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	json.NewEncoder(w).Encode(res)

	err = cmdGoVer.Run()

	if err != nil {
		log.Println("Error running command")
	}

}

func ClientPulling(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit ClientPulling Post Endpoint.")

	var req models.ClientPullingRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.ClientPullingResponse{CrocCodePhrase: "wack"}

	crocExecutable, _ := exec.LookPath("croc")
	cmdGoVer := &exec.Cmd{
		Path:   crocExecutable,
		Args:   []string{crocExecutable, "--version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = cmdGoVer.Run()

	if err != nil {
		log.Println("Error running command")
	}

	json.NewEncoder(w).Encode(res)
}
