// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"

	// Swagger docs.
	_ "github.com/TcMits/ent-clean-template/docs"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface) {
}
