package main

import (
	"github.com/gin-gonic/gin"

	userRepo "github.com/napapol-dev-group/poc-golang/app/repositories/user"

	userUseCase "github.com/napapol-dev-group/poc-golang/app/use_cases/user"

	httpDelivery "github.com/napapol-dev-group/poc-golang/app/handlers/delivery/http"
)

func main() {

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())

	repo := userRepo.InitRepo(nil)

	useCase := userUseCase.InitUseCase(repo)

	httpDelivery.NewEndpointHTTPHandler(ginEngine, useCase)

	err := ginEngine.Run()
	if err != nil {
		panic(err)
	}
}
