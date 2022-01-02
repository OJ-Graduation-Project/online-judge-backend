package contest

import (
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/fenwick"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"github.com/pkg/math"
)

type ScoreBoard struct {
	ScoreToUser          *treemap.Map
	UserToScore          *hashmap.Map
	UserMaxProblemsScore *hashmap.Map
	F                    fenwick.Fenwick
}

func ReverseOrder(a, b interface{}) int {
	return utils.IntComparator(b, a)
}

func New() *ScoreBoard {
	return &ScoreBoard{
		treemap.NewWith(ReverseOrder),
		hashmap.New(),
		hashmap.New(),
		*fenwick.New(10000)}
}

func (s *ScoreBoard) AddProblemScore(user, problemIndex int) {
	z0, _ := s.UserMaxProblemsScore.Get(user)
	problemScores := z0.([]int)

	increase_in_score := problemScores[problemIndex]
	problemScores[problemIndex] = 0

	z1, _ := s.UserToScore.Get(user)
	cur_score := z1.(int)

	s.F.Update(cur_score, -1)

	z2, _ := s.ScoreToUser.Get(cur_score)
	cur_score_set := z2.(*treeset.Set)

	cur_score_set.Remove(user)
	if cur_score_set.Empty() {
		s.ScoreToUser.Remove(cur_score)
	}

	cur_score += increase_in_score

	_, found := s.ScoreToUser.Get(cur_score)
	if !found {
		s.ScoreToUser.Put(cur_score, treeset.NewWithIntComparator())
	}

	z3, _ := s.ScoreToUser.Get(cur_score)
	user_set := z3.(*treeset.Set)

	user_set.Add(user)
	s.F.Update(cur_score, 1)
	s.UserToScore.Put(user, cur_score)
}

func (s *ScoreBoard) Initialize(l, problemsScore []int) {
	user_set := treeset.NewWithIntComparator()
	for _, x := range l {
		tmp := make([]int, len(problemsScore))
		copy(tmp, problemsScore)
		s.UserMaxProblemsScore.Put(x, tmp)
		s.UserToScore.Put(x, 0)
		user_set.Add(x)
	}
	s.ScoreToUser.Put(0, user_set)
	s.F.Update(0, len(l))
}

func (s *ScoreBoard) DecreaseProblemScore(user, problemIndex, value int) {
	x, _ := s.UserMaxProblemsScore.Get(user)
	y := x.([]int)
	y[problemIndex] -= value
	y[problemIndex] = math.Max(y[problemIndex], 0)
}

func (s *ScoreBoard) Get(start_index, count int) pair.PairList {

	a := make(pair.PairList, 0)

	cur := start_index
	end := start_index + count - 1
	size := s.UserToScore.Size()

	first := true

	for cur <= end && cur <= size {

		start_score := s.F.Find(cur)

		z1, _ := s.ScoreToUser.Get(start_score)
		users := z1.(*treeset.Set)
		it := users.Iterator()

		if first {
			idx := s.F.Suffix(start_score + 1)

			for idx+1 < cur && it.Next() {
				idx++
			}

			first = false
		}

		for it.Next() && cur <= end && cur <= size {
			_, value := it.Index(), it.Value()

			a = append(a, pair.New(start_score, value.(int)))
			cur++
		}

	}

	return a
}
