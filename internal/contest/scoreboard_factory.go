package contest

const (
	SLOW = "slow"
	FAST = "fast"
)

func NewScoreBoard(s string) ScoreBoardInterface {
	if s == FAST {
		return NewFastScoreBoard()
	} else if s == SLOW {
		return NewSlowScoreBoard()
	}
	return nil
}
