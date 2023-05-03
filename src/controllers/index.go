package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitControllers(r *gin.Engine) {
	InitUserController(r)
}
