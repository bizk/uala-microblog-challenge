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

// PostTweet godoc
// @Summary Crear un nuevo tweet
// @Description Crear un nuevo tweet
// @Tags tweets
// @Accept json
// @Produce json
// @Param tweet body PostTweetRequest true "Tweet content"
// @Success 201 {string} string "ok"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /tweets [post]
func (h *TweetHandler) PostTweet(c *gin.Context) {
	var req PostTweetRequest
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

// FollowUser godoc
// @Summary Seguir a un usuario
// @Description Seguir a un usuario
// @Tags follow
// @Accept json
// @Produce json
// @Param follow body FollowRequest true "Follow user"
// @Success 201 {string} string "ok"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /tweets/follow [post]
func (h *TweetHandler) FollowUser(c *gin.Context) {
	var req FollowRequest
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

// GetTimeline godoc
// @Summary Get timeline
// @Description Devuelve el timeline de un usuario
// @Tags timeline
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {array} domain.Tweet "Timeline"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /tweets/timeline [get]
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
