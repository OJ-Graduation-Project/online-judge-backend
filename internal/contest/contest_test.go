package contest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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

func TestCorrectness(t *testing.T) {
	var slow Contest = Contest{}
	var fast Contest = Contest{}

	const LENGTH = 100
	const WRONG_SUBMISSION_COST = 1
	const SCORE_DIFF = 1

	slow.RegisteredUserIds = make([]int, LENGTH)
	slow.ContestProblemIds = make([]int, LENGTH)
	slow.ProblemsScore = make([]int, LENGTH)
	slow.WrongSubmissionCost = WRONG_SUBMISSION_COST

	fast.RegisteredUserIds = make([]int, LENGTH)
	fast.ContestProblemIds = make([]int, LENGTH)
	fast.ProblemsScore = make([]int, LENGTH)
	fast.WrongSubmissionCost = WRONG_SUBMISSION_COST

	for i := 1; i <= LENGTH; i++ {
		slow.RegisteredUserIds[i-1] = i
		slow.ContestProblemIds[i-1] = i
		slow.ProblemsScore[i-1] = i * SCORE_DIFF

		fast.RegisteredUserIds[i-1] = i
		fast.ContestProblemIds[i-1] = i
		fast.ProblemsScore[i-1] = i * SCORE_DIFF

	}

	slow.Start("slow")
	fast.Start("fast")

	for i := 0; i <= 2*LENGTH*LENGTH; i++ {
		op := rand.Intn(2)
		userId := 1 + rand.Intn(LENGTH-1)
		problemId := 1 + rand.Intn(LENGTH)
		switch op {

		case 0:
			slow.AcceptedSubmission(userId, problemId)
			fast.AcceptedSubmission(userId, problemId)

		case 1:
			slow.WrongSubmission(userId, problemId)
			fast.WrongSubmission(userId, problemId)
		}

		if slow.DisplayAllRanks() != fast.DisplayAllRanks() {
			fmt.Println(slow.DisplayAllRanks())
			fmt.Println(fast.DisplayAllRanks())
			t.Fail()
			break
		}
	}
}
