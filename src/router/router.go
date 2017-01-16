package router

import (
	"time"

	"github.com/douzl/dockerfile_ui/src/api"
	"github.com/douzl/dockerfile_ui/src/router/middleware"
	"github.com/douzl/dockerfile_ui/src/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Router(middlewares ...gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(utils.Ginrus(log.StandardLogger(), time.RFC3339Nano, false))
	r.Use(middleware.CORSMiddleware())
	r.Use(middlewares...)

	control := api.InitDockerfileUIControl()
	apiv1 := r.Group("/v1")
	{
		apiv1.GET("/ping", control.Ping)
		apiv1.POST("/dockerfile", control.CreateDockerfile)
	}

	return r
}
