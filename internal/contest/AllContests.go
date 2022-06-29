package contest

import (
	"fmt"
	"log"
	"sync"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type AllCont struct {
	m        *sync.Mutex
	contests []*Contest
}

var instantiated *AllCont
var once sync.Once

func GetInstance() *AllCont {
	once.Do(func() {
		instantiated = &AllCont{m: &sync.Mutex{}}
	})
	return instantiated
}

func (all AllCont) AddContest(contest *Contest) {
	all.m.Lock()
	defer all.m.Unlock()
	instantiated.contests = append(instantiated.contests, contest)
}

func (all AllCont) GetContest(id int) *Contest {
	var contest *Contest
	for i := 0; i < len(instantiated.contests); i++ {
		if instantiated.contests[i].ContestId == id {
			contest = instantiated.contests[i]
		}
	}
	return contest
}

func (all AllCont) GetContestAndStart(contestid int64) {
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error couldn't connect to database")
	}
	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{"_id": contestid}, bson.M{})
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)

	}
	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}

	if len(contests) > 1 {
		fmt.Printf("Error more than one Contest with the same ID")
	}

	var ctstData Contest
	bsonBytes, _ := bson.Marshal(contests[0])
	bson.Unmarshal(bsonBytes, &ctstData)
	ctstData.Start("redis")
	fmt.Println(ctstData.Board)
	instantiated.AddContest(&ctstData)
	
}
