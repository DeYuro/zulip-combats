package service

type state string

const (
	idle         state = `idle`
	fightWaiting state = `fightWaiting`
)

type player = string

type states map[player]state
