package arena

import (
	"errors"
)

type FighterType string

const (
	human FighterType = "human"
	ai    FighterType = "ai"
)

type HeathPoint interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type Fighter[HP HeathPoint] interface {
	getHp() HP
	restoreHp()
	takeDamage(damage HP)
}

//Human value type receiver methods
type Human[HP HeathPoint] struct {
	Hp          HP
	MaxHp       HP
	RestoreStep HP
}

func (h Human[HP]) getHp() HP {
	return h.Hp
}

func (h Human[HP]) restoreHp() {
	if h.MaxHp < (h.Hp + h.RestoreStep) {
		h.Hp += h.RestoreStep
	}

	h.Hp = h.MaxHp
}

func (h Human[HP]) takeDamage(damage HP) {
	if (h.Hp - damage) < 0 {
		h.Hp = 0
	}

	h.Hp -= damage
}

// AI pointer type receiver methods
type AI[HP HeathPoint] struct {
	Hp    HP
	MaxHp HP
}

func (a *AI[HP]) getHp() HP {
	return a.Hp
}

func (a *AI[HP]) restoreHp() {
	a.Hp = a.MaxHp
}

func (a *AI[HP]) takeDamage(damage HP) {
	if (a.Hp - damage) < 0 {
		a.Hp = 0
	}

	a.Hp -= damage
}

func createFighter[HP HeathPoint](fighterType FighterType, maxHp, restoreStep HP) (Fighter[HP], error) {
	switch fighterType {
	case human:
		return &Human[HP]{
			Hp:          maxHp,
			MaxHp:       maxHp,
			RestoreStep: restoreStep,
		}, nil
	case ai:
		return &AI[HP]{
			Hp:    maxHp,
			MaxHp: maxHp,
		}, nil
	}

	return nil, errors.New("wrong type")
}
