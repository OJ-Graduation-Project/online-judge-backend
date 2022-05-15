package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/contest"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
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

type Resp struct {
	TotalNumber int             `json:"totalCount"`
	Response    []ScoreResponse `json:"response"`
}

func ScoreBoardHandler(w http.ResponseWriter, r *http.Request) {

	contestID, _ := strconv.Atoi(mux.Vars(r)["id"])
	limit, _ := strconv.Atoi(mux.Vars(r)["limit"])
	pagenumber, _ := strconv.Atoi(mux.Vars(r)["page"])

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()

	ctst := contest.GetInstance().GetContest(contestID)
	if limit > ctst.Board.Count() {
		limit = ctst.Board.Count()
	}
	ans := ctst.GetRanks((pagenumber-1)*limit+1, limit)
	fmt.Println(ctst.DisplayAllRanks())
	mp := make(map[int]int)

	dbconnection, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Error in DB connection")
		log.Fatal(err)
	}
	defer dbconnection.Cancel()

	var userids []int

	for i := 0; i < ans.Len(); i++ {
		userids = append(userids, ans[i].User)
		mp[ans[i].User] = ans[i].Score
	}
	var response []ScoreResponse
	var resp ScoreResponse

	for i := 0; i < ans.Len(); i++ {
		cursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{
			"_id": ans[i].User,
		}, bson.M{})
		var users []bson.M
		if err = cursor.All(dbconnection.Ctx, &users); err != nil {
			fmt.Println("Error in cursor")
			log.Fatal(err)
		}
		resp.Name = users[0]["firstName"].(string)

		val_int, ok := users[0]["_id"].(int64)
		if !ok {
			val_double := users[0]["_id"].(float64)
			resp.UserId = int(val_double)
		}
		resp.UserId = int(val_int)

		resp.Score = ans[i].Score
		response = append(response, resp)

	}

	/*cursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{
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

	for i := 0; i < len(users); i++ {
		resp.Name = users[i]["firstName"].(string)
		resp.UserId = int(users[i]["userId"].(float64))
		resp.Score = mp[int(users[i]["userId"].(float64))]
		response = append(response, resp)
	}
	*/
	fmt.Println(response)
	var respWithTotalCount Resp
	respWithTotalCount.TotalNumber = ctst.Board.Count()
	respWithTotalCount.Response = response

	json.NewEncoder(w).Encode(&respWithTotalCount)

}
