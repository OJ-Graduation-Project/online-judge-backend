package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/contest"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type ScoreRequest struct {
	ContestID int `json:"contestid"`
	Page      int `json:"page"`
}

type ScoreResponse struct {
	Name   string `json:"firstName"`
	UserId int    `json:"userid"`
	Score  int    `json:"score"`
}

type UserD struct {
}

func ScoreBoardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var scorereq ScoreRequest
	err := decoder.Decode(&scorereq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctst := contest.GetInstance().GetContest(scorereq.ContestID)

	ans := ctst.GetRanks(1, ctst.Board.UserToScore.Size())
	mp := make(map[int]int)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	var userids []int

	for i := 0; i < ans.Len(); i++ {
		userids = append(userids, ans[i].User)
		mp[ans[i].User] = ans[i].Score
	}

	cursor, err := dbconnection.Query("OJ_DB", "users", bson.M{
		"userId": bson.M{
			"$in": userids,
		},
	}, bson.M{})

	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var users []bson.M
	if err = cursor.All(dbconnection.Ctx, &users); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}

	var response []ScoreResponse
	var resp ScoreResponse
	for i := 0; i < len(users); i++ {
		resp.Name = users[i]["firstName"].(string)
		resp.UserId = int(users[i]["userId"].(int32))
		resp.Score = mp[int(users[i]["userId"].(int32))]
		response = append(response, resp)
	}

	fmt.Println(response)

	json.NewEncoder(w).Encode(&response)

}