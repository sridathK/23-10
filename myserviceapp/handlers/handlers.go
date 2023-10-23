package handlers

import (
	"fmt"
	"myserviceapp/auth"
	"myserviceapp/middlewear"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth) *gin.Engine {

	r := gin.New()
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.GET("/check", check)
	return r

}

func check(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
	default:
		c.JSON(http.StatusBadGateway, gin.H{"msg": http.StatusOK})
	}
}
