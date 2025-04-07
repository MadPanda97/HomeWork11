package server

import (
	"errors"
	"internet-store/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.userService.UpdateUser(ctx, &entity.UpdateUserRequest{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	})
	if err != nil {
		if errors.Is(err, entity.ErrorInvalidParams) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
