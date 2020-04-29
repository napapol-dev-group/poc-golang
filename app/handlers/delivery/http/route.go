package http

import (
	"github.com/gin-gonic/gin"
	userUseCase "github.com/napapol-dev-group/poc-golang/app/use_cases/user"
)

type handler struct {
	UserUseCase userUseCase.UseCase
}

// NewEndpointHTTPHandler setup router
func NewEndpointHTTPHandler(ginEngine *gin.Engine, userUseCase userUseCase.UseCase) {
	handler := handler{
		UserUseCase: userUseCase,
	}
	group := ginEngine.Group("/user")
	{
		group.POST("/register", handler.Register)
	}
}
