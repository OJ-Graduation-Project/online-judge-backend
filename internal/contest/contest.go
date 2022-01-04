package contest

import (
	"fmt"
	"strings"
	"time"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/emirpasic/gods/maps/hashmap"
)

type Contest struct {
	ContestId               int       `json:"contestId,omitempty"`
	ContestName             string    `json:"contestName,omitempty"`
	StartDate               time.Time `json:"startDate,omitempty"`
	StartTime               string    `json:"startTime,omitempty"`
	Duration                string    `json:"duration,omitempty"`
	NumberOfRegisteredUsers int       `json:"numberOfRegisteredUsers,omitempty"`
	ContestProblemIds       []int     `json:"contestProblemsId,omitempty"`
	RegisteredUserIds       []int     `json:"registeredUsersId,omitempty"`
	ProblemsScore           []int     `json:"problemsScore,omitempty"`
	WrongSubmissionCost     int       `json:"wrongSubmissionCost,omitempty"`
	Board                   *ScoreBoard
	ProblemIdToIndex        *hashmap.Map
}

func (c *Contest) Start() {
	c.ProblemIdToIndex = hashmap.New()

	for i, v := range c.ContestProblemIds {
		c.ProblemIdToIndex.Put(v, i)
	}

	c.Board = New()
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
	for i, p := range a {
		j := start + i
		result.WriteString(fmt.Sprintf("\n%v %v %v\n", j, p.User, p.Score))
	}
	return result.String()
}

func (c *Contest) GetAllRanks() pair.PairList {
	return c.GetRanks(1, c.Board.UserToScore.Size())
}

func (c *Contest) DisplayAllRanks() string {
	return c.DisplayRanks(1, c.Board.UserToScore.Size())
}
