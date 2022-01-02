package pair

type Pair struct {
	Score int
	User  int
}

func New(score, user int) *Pair {
	return &Pair{score, user}
}

type PairList []*Pair

func (s PairList) Len() int {
	return len(s)
}
func (s PairList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PairList) Less(i, j int) bool {
	if s[i].Score == s[j].Score {
		return s[i].User < s[j].User
	}
	return s[i].Score > s[j].Score
}
