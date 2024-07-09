package models

import "sort"

type Leaderboard struct {
	ID    string
	Items []LeaderboardItem
}

type LeaderboardItem struct {
	Owner string `bson:"owner" json:"owner"`
	Index int    `bson:"index" json:"index"`
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
		return leaderboard.Items[i].Index < leaderboard.Items[j].Index
	})
	return leaderboard.Items
}
