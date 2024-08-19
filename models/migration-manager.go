package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(Leaderboard{})
	db.AutoMigrate(LeaderboardItem{})
	db.AutoMigrate(User{})
	db.AutoMigrate(Token{})
}
