package router

import (
    "github.com/gin-gonic/gin"
    "github.com/sxd0/SweetTweet/internal/api/handler"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    auth := handler.NewAuthHandler()
    api := r.Group("/api")
    {
        api.POST("/register", auth.Register)
        api.POST("/login", auth.Login)
        api.POST("/me", auth.Me)
    }

    return r
}
