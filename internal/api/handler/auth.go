package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sxd0/SweetTweet/internal/api/grpcclient"
    "github.com/sxd0/SweetTweet/proto/authpb"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    grpcResp, err := grpcclient.ConnectAuth().Register(c, &authpb.RegisterRequest{
        Email:    req.Email,
        Password: req.Password,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": grpcResp.Message})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    grpcResp, err := grpcclient.ConnectAuth().Login(c, &authpb.LoginRequest{
        Email:    req.Email,
        Password: req.Password,
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": grpcResp.Token})
}

func (h *AuthHandler) Me(c *gin.Context) {
    var req struct {
        Token string `json:"token"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    grpcResp, err := grpcclient.ConnectAuth().Me(c, &authpb.MeRequest{
        Token: req.Token,
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user_id": grpcResp.UserId,
        "email":   grpcResp.Email,
    })
}
