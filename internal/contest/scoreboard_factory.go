package contest

const (
	SLOW = "slow"
	FAST = "fast"
)

func NewScoreBoard(s string) ScoreBoardInterface {
	if s == SLOW {
		return NewFastScoreBoard()
	} else if s == FAST {
		return NewSlowScoreBoard()
	}
	return nil
}
