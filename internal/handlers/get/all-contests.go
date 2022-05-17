package get

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllContests(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}

	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)
	fmt.Println(util.FETCH_ALL_CONTESTS)

	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	fmt.Println(util.RETURNING_ALL_CONTESTS)
	json.NewEncoder(w).Encode(&contests)

}
