package handler

type responseSchema struct {
	GameID  int   `json:"game_id"`
	Numbers []int `json:"numbers"`
}
