package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
	"github.com/gorilla/mux"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	problemID := mux.Vars(r)["problemID"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", problemID)
	
	defer r.Body.Close()
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body from submission: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	
	var submissionRequest requests.SubmissionRequest
	err = json.Unmarshal(body, &submissionRequest)
	if err != nil {
		fmt.Println("Error unmarshalling body from submission: ", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// add submission to database
	// validate submission
	// run submission against input and compare with output

}


