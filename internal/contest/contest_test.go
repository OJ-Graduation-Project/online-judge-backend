package contest

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	// "time"
)

func TestContest(t *testing.T) {

	file, err := ioutil.ReadFile("../../config/contest.json")
	if err != nil {
		t.Error(err)
	}
	var contest Contest
	err = json.Unmarshal(file, &contest)
	if err != nil {
		t.Error(err)
	}

	
	
	contest.Start()

	contest.AcceptedSubmission(1629, 371)
	contest.WrongSubmission(265, 371)
	contest.WrongSubmission(265, 371)
	contest.WrongSubmission(265, 371)
	contest.AcceptedSubmission(265, 371)
	contest.AcceptedSubmission(11, 371)
	t.Log(contest.DisplayAllRanks())
}
