package handlers

import (
	"crocStuff/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func ClientPushing(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit ClientPushing Post Endpoint.")

	var clientInformation models.Request

	err := json.NewDecoder(r.Body).Decode(&clientInformation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.Response{Status: "success"}

	crocExecutable, _ := exec.LookPath("croc")
	cmdGoCroc := &exec.Cmd{
		Path:   crocExecutable,
		Args:   []string{crocExecutable, clientInformation.CrocCodePhrase},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	json.NewEncoder(w).Encode(response)

	defer Execute(cmdGoCroc)
}

func ClientPulling(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit ClientPulling Post Endpoint.")

	var req models.Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.Response{CrocCodePhrase: "wack-abce"}

	crocExecutable, _ := exec.LookPath("croc")
	cmdGoCroc := &exec.Cmd{
		Path:   crocExecutable,
		Args:   []string{crocExecutable, "send", "--code", "wack-abce", "test.txt"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	json.NewEncoder(w).Encode(res)

	defer Execute(cmdGoCroc)
}

func Execute(cmdGoCroc *exec.Cmd) {
	err := cmdGoCroc.Run()

	if err != nil {
		log.Println("Error running command")
	}
}

//The function should accept a request and response writer.
func TestFlow(w http.ResponseWriter, r *http.Request) {

	//logging statement to verify that the endpoint is being hit.
	log.Println("Testing endpoint hit.")

	//There needs to be a struct to store the decoded request body in.
	var clientInfo models.Request

	//A response needs to be instantiated with a default message.
	var response models.Response

	//The request needs to be decoded, and a response needs to be instantiated.
	err := json.NewDecoder(r.Body).Decode(&clientInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//The struct is passed to the verify function that compares the passed values against the database values.
	isValid := Verify(clientInfo)

	//If the data is valid, return a response saying so
	if isValid {
		response.Status = "Valid"
	} else {
		response.Status = "Invalid"
	}

	json.NewEncoder(w).Encode(response)

}

//Verify pulls information from a database and then returns true if the fields match up as expected.
//Can be modified to return a specific error if you wish to know exactly what is wrong.
func Verify(information models.Request) (valid bool) {

	db, err := sql.Open("mysql", "root:Th!3f3ry@tcp(127.0.0.1:3306)/crocTesting")

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	results, err := db.Query("SELECT SC_API_Key, O_API_Key, CO_API_Key FROM userInfo")
	if err != nil {
		log.Println(err)
	}

	for results.Next() {
		var userInfo models.UserInfo

		err = results.Scan(&userInfo.ServerClientAPIKey, &userInfo.OragnizationAPIKey, &userInfo.ClientOrganizationKey)
		if err != nil {
			log.Println(err)
		}

		if(information.ClientOrganizationKey != userInfo.ClientOrganizationKey){
			log.Println("CO_API_KEY Mismatch")
			break
		}
		if(information.ServerClientAPIKey != userInfo.ServerClientAPIKey){
			log.Println("SC_API_KEY Mismatch")
			break
		}
		if(information.OragnizationAPIKey != userInfo.OragnizationAPIKey){
			log.Println("CO_API_KEY Mismatch")
			break
		}

		valid=true
	}

	return valid
}
