package contest

import (
	"strconv"

	// "sort"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/datastructures/pair"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/redis_pool"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/math"
)

var scoreboardMap string = "scoreboardMap"
var problemsMap string = "problemsMap"

type ScoreBoardRedis struct {
	RedisClient   redis.Conn
	problemsScore pair.PairList
}

func NewScoreBoardRedis(problemsScore pair.PairList) *ScoreBoardRedis {
	return &ScoreBoardRedis{redis_pool.NewPool().Get(), problemsScore}
}

func (s *ScoreBoardRedis) AddProblemScore(userId, problemId int) {

	key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemId)
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

	for _, userId := range userIds {
		s.RedisClient.Do("ZADD", scoreboardMap, 0, strconv.Itoa(userId))
		for _, problemData := range s.problemsScore {
			problemId, maxProblemScore := problemData.Id, problemData.Score
			key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemId)
			s.RedisClient.Do("HSET", problemsMap, key, strconv.Itoa(maxProblemScore))
		}
	}

}

func (s *ScoreBoardRedis) DecreaseProblemScore(userId, problemId, value int) {
	key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemId)

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
	size, _ := redis.Int(s.RedisClient.Do("ZCOUNT", scoreboardMap, "-inf", "inf"))
	return size
}

func (s *ScoreBoardRedis) Register(userId int) {
	s.RedisClient.Do("ZADD", scoreboardMap, 0, strconv.Itoa(userId))
	for _, problemData := range s.problemsScore {
		problemId, maxProblemScore := problemData.Id, problemData.Score
		key := strconv.Itoa(userId) + ", " + strconv.Itoa(problemId)
		s.RedisClient.Do("HSET", problemsMap, key, strconv.Itoa(maxProblemScore))
	}
}

func (s *ScoreBoardRedis) IsRegistered(userId int) bool {
	key := strconv.Itoa(userId) + ", " + strconv.Itoa(s.problemsScore[0].Id)
	exist, _ := redis.Int(s.RedisClient.Do("HEXISTS", problemsMap, key))
	return exist == 1
}
