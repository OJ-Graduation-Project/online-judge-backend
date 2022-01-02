package contest

import (
	"sort"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
)

type ScoreBoard_ struct {
	Board pair.PairList
}

func NewScoreBoard_() *ScoreBoard_ {
	return &ScoreBoard_{make([]*pair.Pair, 0)}
}

func (s *ScoreBoard_) UpdateScore(user, increase_in_score int) {
	for i := 0; i < len(s.Board); i++ {
		if s.Board[i].User == user {
			s.Board[i].Score += increase_in_score
			break
		}
	}
	sort.Sort(s.Board)
}

func (s *ScoreBoard_) Initialize(l []int) {
	for _, v := range l {
		s.Board = append(s.Board, pair.New(0, v))
	}
}

func (s *ScoreBoard_) Get(start_index, count int) []int {
	sort.Sort(s.Board)

	idx := 0
	x := make([]int, count)

	for i := start_index - 1; i < len(s.Board) && idx < count; i++ {
		x[idx] = s.Board[i].User
		idx++
	}

	return x
}
