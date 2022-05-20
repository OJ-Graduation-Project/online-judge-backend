package contest

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/emirpasic/gods/maps/hashmap"
)

type Contest struct {
	ContestId               int       `bson:"_id,omitempty" json:"contestId,omitempty"`
	ContestName             string    `bson:"contestname,omitempty" json:"contestName,omitempty"`
	StartDate               time.Time `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartTime               string    `bson:"startTime,omitempty" json:"startTime,omitempty"`
	Duration                string    `bson:"duration,omitempty" json:"duration,omitempty"`
	NumberOfRegisteredUsers int       `bson:"numberOfRegisteredUsers,omitempty" json:"numberOfRegisteredUsers,omitempty"`
	ContestProblemIds       []int     `bson:"contest_problemset,omitempty" json:"contestProblemId,omitempty"`
	RegisteredUserIds       []int     `bson:"registeredUsersId,omitempty" json:"registeredUserId,omitempty"`
	ProblemsScore           []int     `bson:"problemsscore,omitempty" json:"problemsScore,omitempty"`
	WrongSubmissionCost     int       `bson:"wrongSubmissionCost,omitempty" json:"wrongSubmissionCost,omitempty"`
	Board                   *ScoreBoardRedis
	ProblemIdToIndex        *hashmap.Map
}

func (c *Contest) Start(scoreBoardType string) {
	problemsScore := make(pair.PairList, len(c.ProblemsScore))

	for i, _ := range c.ContestProblemIds {
		problemsScore[i] = pair.New(c.ProblemsScore[i], c.ContestProblemIds[i])
	}

	c.Board = NewScoreBoardRedis(problemsScore)

}

func (c *Contest) AcceptedSubmission(userId, problemId int) {
	c.Register(userId)
	c.Board.AddProblemScore(userId, problemId)
}

func (c *Contest) WrongSubmission(userId, problemId int) {
	c.Register(userId)
	c.Board.DecreaseProblemScore(userId, problemId, c.WrongSubmissionCost)
}

func (c *Contest) GetRanks(start, count int) pair.PairList {
	return c.Board.Get(start, count)
}

func (c *Contest) DisplayRanks(start, count int) string {
	a := c.GetRanks(start, count)
	var result strings.Builder
	sort.Sort(a) // This line can be removed if we don't care about how will we break ties
	for i, p := range a {
		j := start + i
		result.WriteString(fmt.Sprintf("\n%v %v %v\n", j, p.Id, p.Score))
	}
	return result.String()
}

func (c *Contest) GetAllRanks() pair.PairList {
	return c.GetRanks(1, c.Board.Count())
}

func (c *Contest) DisplayAllRanks() string {
	return c.DisplayRanks(1, c.Board.Count())
}
func (c *Contest) Register(userId int) {
	if !c.Board.IsRegistered(userId) {
		c.Board.Register(userId)
	}
}
