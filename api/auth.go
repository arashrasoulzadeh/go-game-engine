package api

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AuthRoutes(engine *gin.Engine, db *gorm.DB, cn *gochannel.GoChannel) {
	v1 := engine.Group("/v1/auth")
	{
		v1.GET("guest", V1AuthGuest(db, cn))
		v1.GET("profile", V1AuthProfile(db, cn))
	}
}

func V1AuthGuest(db *gorm.DB, cn *gochannel.GoChannel) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var token models.Token

		user.Name = "guest"

		db.Create(&user)
		token.UserID = user.ID

		password := []byte(strconv.Itoa(int(user.ID)))

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		token.Token = string(hashedPassword)
		db.Create(&token)

		c.JSON(200, gin.H{
			"token": token.Token,
		})

		models.GetBus().Bus <- "guest created!"

	}
}

func V1AuthProfile(db *gorm.DB, cn *gochannel.GoChannel) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var user models.User
		var token models.Token

		if c.Request.Header.Get("Authorization") == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		db.First(&token, "token = ?", c.Request.Header.Get("Authorization"))

		if token.Token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User

		db.First(&user, "id = ?", token.UserID)

		c.JSON(200, gin.H{
			"token": token.Token,
			"user":  user,
		})
	}
}
