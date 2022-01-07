package contest

import (
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
)

type ScoreBoardInterface interface {
	AddProblemScore(user, problemIndex int)
	Initialize(userIds, problemsScore []int)
	DecreaseProblemScore(user, problemIndex, value int)
	Get(start_index, count int) pair.PairList
	Count() int
}
