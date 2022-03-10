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
	ContestId               int       `bson:"contestId,omitempty" json:"contestId,omitempty"`
	ContestName             string    `bson:"contestName,omitempty" json:"contestName,omitempty"`
	StartDate               time.Time `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartTime               string    `bson:"startTime,omitempty" json:"startTime,omitempty"`
	Duration                string    `bson:"duration,omitempty" json:"duration,omitempty"`
	NumberOfRegisteredUsers int       `bson:"numberOfRegisteredUsers,omitempty" json:"numberOfRegisteredUsers,omitempty"`
	ContestProblemIds       []int     `bson:"contestProblemId,omitempty" json:"contestProblemId,omitempty"`
	RegisteredUserIds       []int     `bson:"registeredUserId,omitempty" json:"registeredUserId,omitempty"`
	ProblemsScore           []int     `bson:"problemsScore,omitempty" json:"problemsScore,omitempty"`
	WrongSubmissionCost     int       `bson:"wrongSubmissionCost,omitempty" json:"wrongSubmissionCost,omitempty"`
	Board                   ScoreBoardInterface
	ProblemIdToIndex        *hashmap.Map
}

func (c *Contest) Start(scoreBoardType string) {

	c.ProblemIdToIndex = hashmap.New()

	for i, v := range c.ContestProblemIds {
		c.ProblemIdToIndex.Put(v, i)
	}

	c.Board = NewScoreBoard(scoreBoardType)
	c.Board.Initialize(c.RegisteredUserIds, c.ProblemsScore)
}

func (c *Contest) AcceptedSubmission(userId, problemId int) {
	z0, _ := c.ProblemIdToIndex.Get(problemId)
	problemIndex := z0.(int)
	c.Board.AddProblemScore(userId, problemIndex)
}

func (c *Contest) WrongSubmission(userId, problemId int) {
	z0, _ := c.ProblemIdToIndex.Get(problemId)
	problemIndex := z0.(int)
	c.Board.DecreaseProblemScore(userId, problemIndex, c.WrongSubmissionCost)
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
		result.WriteString(fmt.Sprintf("\n%v %v %v\n", j, p.User, p.Score))
	}
	return result.String()
}

func (c *Contest) GetAllRanks() pair.PairList {
	return c.GetRanks(1, c.Board.Count())
}

func (c *Contest) DisplayAllRanks() string {
	return c.DisplayRanks(1, c.Board.Count())
}
