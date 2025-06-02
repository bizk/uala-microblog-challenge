package main

import (
	"log"
	"uala-microblog-challenge/internal/infrastructure/memory"
	"uala-microblog-challenge/internal/interfaces/http"
	"uala-microblog-challenge/internal/usecase"

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

// TODO agregar errores tipados
// TODO agregar mejores logs
func main() {
	router := gin.Default()

	// Aca podriamos cargar distintas bases de datos segun una variable de ambiente por ejemplo
	userRepo := memory.NewInMemoryUserRepository()
	tweetRepo := memory.NewInMemoryTweetRepository()

	// Usamos inyeccion de dependencias por que nos permite separar la implementacion de la interfaz. El principal beneficio de esto es poder testear y modificar con facilidad sin afectar el codigo ya implementado (osea desacoplar)
	// PD: Es mi patron favorito :p
	newTweet := usecase.NewPostTweetUsecase(tweetRepo, userRepo)
	newFollowUser := usecase.NewFollowUserUsecase(userRepo)
	newGetTimeline := usecase.NewGetTimelineUsecase(tweetRepo, userRepo)

	tweetHandler := http.NewTweetHandler(newTweet, newFollowUser, newGetTimeline)

	createUser := usecase.NewCreateUserUsecase(userRepo)
	getUser := usecase.NewGetUserUsecase(userRepo)
	userHandler := http.NewUserHandler(createUser, getUser)

	v1 := router.Group("/api/v1")
	tweetGroup := v1.Group("/tweets")
	{
		tweetGroup.POST("", tweetHandler.PostTweet)
		tweetGroup.POST("/follow", tweetHandler.FollowUser)
		tweetGroup.GET("/timeline", tweetHandler.GetTimeline)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Aca podriamos definir TLS / HTTPS
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
