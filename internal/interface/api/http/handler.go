package http

import (
	"net/http"
	"time"

	"github.com/Coke15/AlphaWave-BackEnd/internal/domain/service"
	v1 "github.com/Coke15/AlphaWave-BackEnd/internal/interface/api/http/v1"
	"github.com/Coke15/AlphaWave-BackEnd/pkg/auth/manager"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service         *service.Service
	JWTManager      *manager.JWTManager
	refreshTokenTTL time.Duration
}

func NewHandler(service *service.Service, JWTManager *manager.JWTManager, refreshTokenTTL time.Duration) *Handler {
	return &Handler{
		service:         service,
		JWTManager:      JWTManager,
		refreshTokenTTL: refreshTokenTTL,
	}

}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode("debug")
	r.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	h.initApi(r)
	return r
}

func (h *Handler) initApi(r *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service, h.JWTManager, h.refreshTokenTTL)

	api := r.Group("/api")
	{
		handlerV1.InitRoutes(api)
	}
}
