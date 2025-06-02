package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewTweetHandler(t NewTweet, f Follow, tl Timeline) *TweetHandler {
	return &TweetHandler{
		newTweet: t,
		follow:   f,
		timeline: tl,
	}
}

// ---- New Tweet ----

func (h *TweetHandler) PostTweet(c *gin.Context) {
	var req postTweetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.newTweet.PostTweet(req.UserID, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create tweet"})
		return
	}
	c.Status(http.StatusCreated)
}

// ---- Follow ----

func (h *TweetHandler) FollowUser(c *gin.Context) {
	var req followRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.follow.Follow(req.FollowerID, req.FollowedID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not follow user"})
		return
	}
	c.Status(http.StatusCreated)
}

// ---- Timeline ----

func (h *TweetHandler) GetTimeline(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	tweets, err := h.timeline.GetTimeline(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch timeline"})
		return
	}
	c.JSON(http.StatusOK, tweets)
}
