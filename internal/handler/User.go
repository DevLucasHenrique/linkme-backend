package handler

import (
	"net/http"

	"github.com/DevLucasHenrique/mymock-backend/internal/entity"
	"github.com/DevLucasHenrique/mymock-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (uh *UserHandler) ListAllUsers(ctx *gin.Context) {
	users, err := uh.UserUsecase.ListAllUsers(ctx)
	if err!=nil {
		r := entity.Response{
			Message: "Can't return users, try again latter",
			Code: 500,
		}
		ctx.JSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, users)
}