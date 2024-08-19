package models

import "github.com/gin-gonic/gin"

type Server struct {
	Port     string
	Hostname string
	Engine   *gin.Engine
}
