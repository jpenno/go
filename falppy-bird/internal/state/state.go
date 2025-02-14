package state

type APP_STATE int

const (
	START APP_STATE = iota
	PLAY
	PAUSE
	GAME_OVER
	QUIT
)
