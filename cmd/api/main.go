package main

import (
	"log"
	"uala-microblog-challenge/internal/interfaces/http"

	_ "uala-microblog-challenge/docs" // generated docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Uala Microblog Challenge
// @version         1.0
// @description     Servidor HTTP para el microblog de Uala.

// @contact.name   Carlos Santiago Yanzon
// @contact.url    https://github.com/bizk

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api
func main() {
	router := gin.Default()

	tweetHandler := http.NewTweetHandler(nil, nil, nil) // TODO: Implementar los casos de uso

	v1 := router.Group("/api/v1")
	{
		v1.POST("/tweets", tweetHandler.PostTweet)
		v1.POST("/follow", tweetHandler.FollowUser)
		v1.GET("/timeline", tweetHandler.GetTimeline)
	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
