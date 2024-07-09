package models

import "sort"

type Leaderboard struct {
	ID    string
	Items []LeaderboardItem
}

type LeaderboardItem struct {
	Owner UserID `bson:"owner" json:"owner"`
	Score int    `bson:"score" json:"score"`
}

func CreateLeaderboard() *Leaderboard {
	return &Leaderboard{
		ID:    "-1",
		Items: []LeaderboardItem{},
	}
}

func (leaderboard *Leaderboard) Init() {

}

func (leaderboard *Leaderboard) Add(item LeaderboardItem) {
	leaderboard.Items = append(leaderboard.Items, item)
}

func (leaderboard *Leaderboard) GetID() string {
	return leaderboard.ID
}

func (leaderboard *Leaderboard) SetID(id string) {
	leaderboard.ID = id
}

func (leaderboard *Leaderboard) Sorted() []LeaderboardItem {
	sort.Slice(leaderboard.Items, func(i, j int) bool {
		return leaderboard.Items[i].Score < leaderboard.Items[j].Score
	})
	return leaderboard.Items
}
