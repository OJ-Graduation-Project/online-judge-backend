package contest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestContest(t *testing.T) {

	file, err := ioutil.ReadFile("../../scripts/database/data/contests.json")
	if err != nil {
		t.Error(err)
	}
	var x [1]Contest
	err = json.Unmarshal(file, &x)
	if err != nil {
		t.Error(err)
	}
	var y [1]Contest
	err = json.Unmarshal(file, &y)
	if err != nil {
		t.Error(err)
	}

	fastContest := x[0]
	fastContest.Start("fast")

	slowContest := y[0]
	slowContest.Start("slow")

	fastContest.AcceptedSubmission(1629, 371)
	fastContest.WrongSubmission(265, 371)
	fastContest.WrongSubmission(265, 371)
	fastContest.WrongSubmission(265, 371)
	fastContest.AcceptedSubmission(265, 371)
	fastContest.AcceptedSubmission(11, 371)

	slowContest.AcceptedSubmission(1629, 371)
	slowContest.WrongSubmission(265, 371)
	slowContest.WrongSubmission(265, 371)
	slowContest.WrongSubmission(265, 371)
	slowContest.AcceptedSubmission(265, 371)
	slowContest.AcceptedSubmission(11, 371)

	fmt.Println("Fast Contest Ranks")
	fastRanks := fastContest.DisplayAllRanks()
	fmt.Println(fastRanks)
	fmt.Println("Slow Contest Ranks")
	slowRanks := slowContest.DisplayAllRanks()
	fmt.Println(slowRanks)
	if fastRanks != slowRanks {
		t.Fail()
	}
}
