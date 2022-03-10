package contest

import (
	"strconv"

	// "sort"
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/redis_pool"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/math"
)

var scoreboardMap string
var problemsMap string

type ScoreBoardRedis struct {
	Board                pair.PairList
	UserMaxProblemsScore *hashmap.Map
	RedisClient 	  	 redis.Conn
}

func NewScoreBoardRedis() *ScoreBoardRedis {
	return &ScoreBoardRedis{make([]*pair.Pair, 0), hashmap.New(), redis_pool.NewPool().Get(),}
}

func (s *ScoreBoardRedis) AddProblemScore(userId, problemIndex int) {
	
	key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemIndex)
	reply, err := s.RedisClient.Do("HGET", problemsMap, key)
	if err != nil {
		panic(err.Error())
	}
	increase_in_score, err := redis.Int(reply, err)
	if err != nil {
		panic(err.Error())
	}
	increase_in_score = math.Max(increase_in_score, 0)

	_, err = s.RedisClient.Do("HSET", problemsMap, key, "0")
	if err != nil {
		panic(err.Error())
	}
	_, err = s.RedisClient.Do("ZINCRBY", scoreboardMap, strconv.Itoa(increase_in_score), strconv.Itoa(userId))
	if err != nil {
		panic(err.Error())
	}
	
}

func (s *ScoreBoardRedis) Initialize(userIds, problemsScore []int) {
	scoreboardMap = fmt.Sprintf("scoreboard%p", s)
	problemsMap = fmt.Sprintf("problemsscores%p", s)
	fmt.Println(scoreboardMap, problemsMap)
	
	s.RedisClient.Do("ZREMRANGEBYSCORE", scoreboardMap, "-inf", "inf")

	for _, userId := range userIds {
		tmp := make([]int, len(problemsScore))
		copy(tmp, problemsScore)
		s.UserMaxProblemsScore.Put(userId, tmp)
		s.Board = append(s.Board, pair.New(0, userId))
	}

	for _, userId := range userIds {
		s.RedisClient.Do("ZADD", scoreboardMap, 0, strconv.Itoa(userId))
		for problemIndex, maxProblemScore := range problemsScore {
			key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemIndex)
			s.RedisClient.Do("HSET", problemsMap, key, strconv.Itoa(maxProblemScore))
		}
	}
	
}

func (s *ScoreBoardRedis) DecreaseProblemScore(userId, problemIndex, value int) {
	key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemIndex)
	
	_, err := s.RedisClient.Do("HINCRBY", problemsMap, key, strconv.Itoa(-value))
	if err != nil {
		panic(err.Error())
	}
	
}

func (s *ScoreBoardRedis) Get(startIndex, count int) pair.PairList {
	
	startIndex--
	endIndex := startIndex + count - 1
	value, _ := redis.Ints(s.RedisClient.Do("ZREVRANGE", scoreboardMap, strconv.Itoa(startIndex), strconv.Itoa(endIndex), "WITHSCORES"))
	x := make(pair.PairList, count)

	for i, j := 0, 0; i < len(value); i, j = i+2, j+1 {
		x[j] = pair.New(value[i+1], value[i])
	}
	
	return x
}

func (s *ScoreBoardRedis) Count() int {
	return s.UserMaxProblemsScore.Size()
}
