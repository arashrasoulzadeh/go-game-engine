package api

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunServer(server *models.Server, db *gorm.DB, c *gochannel.GoChannel) {
	initRoutes(server.Engine, db, c)
	err := server.Engine.Run(server.Hostname + ":" + server.Port)
	if err != nil {
		panic(err.Error())
		return
	}
}

func initRoutes(engine *gin.Engine, db *gorm.DB, c *gochannel.GoChannel) {
	LeaderBoardRoutes(engine, db, c)
	AuthRoutes(engine, db, c)
	AgentRoutes(engine, db, c)
}
