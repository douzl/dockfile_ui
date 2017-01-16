package api

import (
	"github.com/douzl/dockerfile_ui/src/service"
	"github.com/douzl/dockerfile_ui/src/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type DockerfileUIControl struct {
	Dockerfile *service.Dockfile
}

func InitDockerfileUIControl() *DockerfileUIControl {
	return &DockerfileUIControl{
		Dockerfile: service.NewDockerfile(),
	}
}

func (dc *DockerfileUIControl) Ping(ctx *gin.Context) {
	utils.Ok(ctx, "success")
}

func (dc *DockerfileUIControl) CreateDockerfile(ctx *gin.Context) {
	if err := ctx.BindJSON(&dc.Dockerfile); err != nil {
		log.Error("invalid param")
		utils.ErrorResponse(ctx, err)
		return
	}

	if err := dc.Dockerfile.WriteDockerfile(); err != nil {
		log.Error(err)
		utils.ErrorResponse(ctx, err)
		return
	}

	utils.Create(ctx, "success")
}
