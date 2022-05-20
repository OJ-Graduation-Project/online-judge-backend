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

	fmt.Println()
	fmt.Println(util.DECODE_TOPIC)
	err := decoder.Decode(&topicProblems)
	if err != nil {
		fmt.Println(util.DECODE_TOPIC_FAILED)
		log.Fatal(err)
		return
	}

	fmt.Println(util.DECODE_TOPIC_SUCCESS)

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	fmt.Println(util.PING_DATABASE)
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println(util.PING)
		log.Fatal(err)
		return
	}

	fmt.Println(util.FETCHING_PROBLEMS_FROM_TOPIC + topicProblems.Name)
	 query := bson.M{"topic": bson.M{"$in": bson.A{topicProblems.Name}}}
	// desiredProblems := QueryToCheckResults(dbconnection, util.PROBLEMS_COLLECTION, query)
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, query, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var desiredProblems []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &desiredProblems); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	if len(desiredProblems) == 0 {
		fmt.Println(util.EMPTY_TOPIC_PROBLEMS)
		json.NewEncoder(w).Encode(bson.M{"message": util.EMPTY_TOPIC_PROBLEMS})
		return
	}

	fmt.Println(util.RETURNING_DESIRED_PROBLEMS)
	json.NewEncoder(w).Encode(&desiredProblems)

}
