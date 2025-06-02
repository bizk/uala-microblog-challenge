package http

// Definimos los casos de uso para mantenerlos aislados del resto del sistema

// A mi en particular me gusta separar los types del resto del codigo y tenerlo todo empaquetado segun.
// Es decir, no tendria todo dentro de http. Seguramente haria algo asi como http/tweet
// Otra alternativa que usaria de tener muchos casos de uso es unificar types y functions, segmentar segun corresponde
// Ej: Juntamos los types de follow y la funcion. Lo ponemos en un solo paquete
// En codebases muy grandess esta bueno por que sin tener mucha idea podes entender de donde sale que
type TweetHandler struct {
	newTweet NewTweet
	follow   Follow
	timeline Timeline
}

// ---- New Tweet ----

type NewTweet interface {
	PostTweet(userID string, content string) error
}

type PostTweetRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required,max=280"` // Seteamos el max de caracteres a 280
}

// ---- Follow ----

type Follow interface {
	Follow(followerID string, followedID string) error
}

type FollowRequest struct {
	FollowerID string `json:"follower_id" binding:"required"`
	FollowedID string `json:"followed_id" binding:"required"`
}

type Timeline interface {
	GetTimeline(userID string) ([]TweetDTO, error)
}

// ---- Tweet DTO ----

type TweetDTO struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Time    string `json:"timestamp"`
}
