package models

import (
	"gorm.io/gorm"
	"sort"
)

type Leaderboard struct {
	gorm.Model
	Items []LeaderboardItem `gorm:"foreignKey:LeaderboardID;references:ID" json:"items" bson:"items"`
}

type LeaderboardItem struct {
	gorm.Model
	LeaderboardID uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"leaderboard_id" bson:"leaderboard_id"`
	Owner         UserID `bson:"owner" json:"owner"`
	Score         int    `bson:"score" json:"score" gorm:"index"`
}

func CreateLeaderboard() *Leaderboard {
	return &Leaderboard{
		Items: []LeaderboardItem{},
	}
}

func (leaderboard *Leaderboard) Init() {

}

func (leaderboard *Leaderboard) Add(item LeaderboardItem) {
	leaderboard.Items = append(leaderboard.Items, item)
}

func (leaderboard *Leaderboard) Sorted() []LeaderboardItem {
	sort.Slice(leaderboard.Items, func(i, j int) bool {
		return leaderboard.Items[i].Score > leaderboard.Items[j].Score
	})
	return leaderboard.Items
}
