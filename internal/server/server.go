package server

import (
	"internet-store/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userService service.UserService
	router      *gin.Engine
	//logger  *logrus.Logger
}

func NewServer(userService service.UserService) *Server {
	return &Server{userService: userService}
}

func (s *Server) SetupRouter() *gin.Engine {
	s.router = gin.New()

	s.router.PATCH("/users", s.UpdateUser)
	s.router.POST("/users", s.UpdateUser)

	return s.router
}
