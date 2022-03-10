package contest

import (
	"fmt"

	"math/rand"
	"testing"
)

func TestCorrectness(t *testing.T) {

	const USERS = 3000
	const REQUESTS = 10 * USERS
	const SUBMIT_THRESHOLD = 30
	const PROBLEMS = 10
	const PAGESIZE = 1
	const MAXMUL = USERS / PAGESIZE
	const WRONG_SUBMISSION_COST = 1
	const SCORE_DIFF = 1
	contestTypes := []string{"fast", "slow", "redis"}
	contests := make([]Contest, len(contestTypes))

	for i, contestType := range contestTypes {
		contests[i] = Contest{}
		prepareContest(&contests[i], USERS, PROBLEMS, WRONG_SUBMISSION_COST, SCORE_DIFF).Start(contestType)
	}

	submissions, fetches := 0, 0
	for i := 0; i <= REQUESTS; i++ {
		if i%500000 == 0 {
			fmt.Println(i)
		}
		if rand.Intn(100) < SUBMIT_THRESHOLD {
			op := rand.Intn(2)
			userId := 1 + rand.Intn(USERS)
			problemId := 1 + rand.Intn(PROBLEMS)
			switch op {
			case 0:
				for _, contest := range contests {
					contest.AcceptedSubmission(userId, problemId)
				}
			case 1:
				for _, contest := range contests {
					contest.WrongSubmission(userId, problemId)
				}
			}
			submissions++
		} else {
			fetches++
			skippedPages := rand.Intn(MAXMUL) + 1
			for _, contest := range contests {
				contest.GetRanks(skippedPages*PAGESIZE, PAGESIZE)
			}
		}
	}
	fmt.Printf("Submissions=%v, Fetches=%v\n", submissions, fetches)
	check(t, contests)
}

func prepareContest(contest *Contest, USERS, PROBLEMS, WRONG_SUBMISSION_COST, SCORE_DIFF int) *Contest {
	contest.RegisteredUserIds = make([]int, USERS)
	contest.ContestProblemIds = make([]int, PROBLEMS)
	contest.ProblemsScore = make([]int, PROBLEMS)
	contest.WrongSubmissionCost = WRONG_SUBMISSION_COST
	for i := 1; i <= USERS; i++ {
		contest.RegisteredUserIds[i-1] = i
	}

	for i := 1; i <= PROBLEMS; i++ {
		contest.ContestProblemIds[i-1] = i
		contest.ProblemsScore[i-1] = i * SCORE_DIFF
	}
	return contest
}
func check(t *testing.T, contests []Contest) {

	for i := 1; i < len(contests); i++ {
		if contests[i].DisplayAllRanks() != contests[i-1].DisplayAllRanks() {
			fmt.Println(i - 1)
			fmt.Println(contests[i-1].DisplayAllRanks())
			fmt.Println("--------------------------------------------------------")
			fmt.Println(i)
			fmt.Println(contests[i].DisplayAllRanks())
			t.Fail()
		}
	}
}

func helper(redis, slow Contest) {

	if slow.DisplayAllRanks() != redis.DisplayAllRanks() {
		fmt.Println("big trouble")
		fmt.Println("Redis")
		fmt.Println(redis.DisplayAllRanks())
		fmt.Println("----------------------")
		fmt.Println("Slow")
		fmt.Println(slow.DisplayAllRanks())
		return
	}
}
