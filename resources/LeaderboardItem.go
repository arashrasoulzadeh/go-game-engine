package resources

type LeaderboardItem struct {
	Score int    `json:"score"`
	Name  string `json:"name"`
}

func (l LeaderboardItem) String() string {
	return l.Name
}
