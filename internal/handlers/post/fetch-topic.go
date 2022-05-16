package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type DisplayTopic struct {
	Name string `json:"topicName,omitempty"`
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var topicProblems DisplayTopic
	err := decoder.Decode(&topicProblems)
	if err != nil {
		fmt.Println("Error couldn't decode problem")
		log.Fatal(err)
		return
	}
	fmt.Println(topicProblems)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		log.Fatal(err)
		return
	}
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println("Error in PING")
		log.Fatal(err)
		return
	}
	query := bson.M{"topic": bson.M{"$in": bson.A{topicProblems.Name}}}

	desiredProblems := QueryToCheckResults(dbconnection, util.PROBLEMS_COLLECTION, query)
	json.NewEncoder(w).Encode(&desiredProblems)

}
