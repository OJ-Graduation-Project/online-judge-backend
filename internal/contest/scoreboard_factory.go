package contest

const (
	SLOW  = "slow"
	FAST  = "fast"
	REDIS = "redis"
)

func NewScoreBoard(s string) ScoreBoardInterface {
	if s == FAST {
		return NewFastScoreBoard()
	} else if s == SLOW {
		return NewSlowScoreBoard()
	} else if s == REDIS {
		return NewScoreBoardRedis()
	}
	return nil
}
