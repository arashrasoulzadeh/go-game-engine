package api

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeaderBoardRoutes(engine *gin.Engine, db *gorm.DB, c *gochannel.GoChannel) {
	v1 := engine.Group("/v1/leaderboard/")
	{
		v1.GET("index", V1LeaderBoardIndex(db))
	}
}

func V1LeaderBoardIndex(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var leaderboard models.Leaderboard
		result := db.Find(&leaderboard, 1)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		var items []models.LeaderboardItem
		resultItems := db.Where("leaderboard_id = ?", 1).Order("score desc").Find(&items)
		if resultItems.Error != nil {
			c.JSON(500, gin.H{"error": resultItems.Error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"data":  leaderboard,
			"items": items,
		})
	}
}
