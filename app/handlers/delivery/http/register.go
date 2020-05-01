package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/napapol-dev-group/poc-golang/app/handlers/delivery/http/models"
	"github.com/napapol-dev-group/poc-golang/app/utils"
	"net/http"
)

func (h *handler) Register(c *gin.Context) {

	var requestData models.RegisterUserRequest
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.UserUseCase.RegisterUser(requestData.Entity())
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	utils.JSONSuccessResponse(c, createdUser)

}
