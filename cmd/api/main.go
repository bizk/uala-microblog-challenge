package main

import (
	"log"
	"uala-microblog-challenge/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	tweetHandler := http.NewTweetHandler(nil, nil, nil) // TODO: Implementar los casos de uso

	router.POST("/tweets", tweetHandler.PostTweet)
	router.POST("/follow", tweetHandler.FollowUser)
	router.GET("/timeline", tweetHandler.GetTimeline)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
