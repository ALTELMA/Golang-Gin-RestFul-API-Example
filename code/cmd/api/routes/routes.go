package routes

import (
	"github.com/gin-gonic/gin"
	"open-market.com/user-api/internal/api"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/member", api.GetMemberById)
	}
}
