package utils

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Leaderboard struct {
	URL  string `json:"url"`
	Hits int    `json:"hits"`
}
