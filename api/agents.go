package api

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var workers []byte

func AgentRoutes(engine *gin.Engine, db *gorm.DB, cn *gochannel.GoChannel) {
	v1 := engine.Group("/v1/agents")
	{
		v1.GET("", V1Agents(db, cn))
	}

}

func V1Agents(db *gorm.DB, cn *gochannel.GoChannel) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"list": agent.GetAgentsPool(),
		})

		models.GetBus().Bus <- "guest created!"

	}
}
