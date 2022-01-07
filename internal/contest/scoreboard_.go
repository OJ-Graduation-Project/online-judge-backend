package contest

import (
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/pkg/math"
	"sort"
)

type ScoreBoard_ struct {
	Board                pair.PairList
	UserMaxProblemsScore *hashmap.Map
}

func NewSlowScoreBoard() *ScoreBoard_ {
	return &ScoreBoard_{make([]*pair.Pair, 0), hashmap.New()}
}

func (s *ScoreBoard_) AddProblemScore(userId, problemIndex int) {
	z0, _ := s.UserMaxProblemsScore.Get(userId)
	problemScores := z0.([]int)

	increase_in_score := problemScores[problemIndex]
	problemScores[problemIndex] = 0

	for i := 0; i < len(s.Board); i++ {
		if s.Board[i].User == userId {
			s.Board[i].Score += increase_in_score
			break
		}
	}
	sort.Sort(s.Board)
}

func (s *ScoreBoard_) Initialize(userIds, problemsScore []int) {
	for _, x := range userIds {
		tmp := make([]int, len(problemsScore))
		copy(tmp, problemsScore)
		s.UserMaxProblemsScore.Put(x, tmp)

		s.Board = append(s.Board, pair.New(0, x))
	}
}

func (s *ScoreBoard_) DecreaseProblemScore(userId, problemIndex, value int) {
	x, _ := s.UserMaxProblemsScore.Get(userId)
	y := x.([]int)
	y[problemIndex] -= value
	y[problemIndex] = math.Max(y[problemIndex], 0)
}

func (s *ScoreBoard_) Get(startIndex, count int) pair.PairList {
	sort.Sort(s.Board)

	idx := 0
	x := make(pair.PairList, count)

	for i := startIndex - 1; i < len(s.Board) && idx < count; i++ {
		x[idx] = s.Board[i]
		idx++
	}

	return x
}

func (s *ScoreBoard_) Count() int {
	return len(s.Board)
}
