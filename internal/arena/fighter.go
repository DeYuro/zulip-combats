package arena

import "golang.org/x/exp/constraints"

type FighterType string

const human FighterType = "human"
const ai FighterType = "ai"

type Fighter[HP constraints.Integer] interface {
	getHP() HP
	restoreHP()
}

type Human[HP constraints.Integer] struct {
	HealthPoint   HP
	MaxHP         HP
	RestoreHpStep HP
}

func (h Human[HP]) getHP() HP {
	return h.getHP()

}

func (h Human) restoreHP() {
	if (h.HealthPoint + h.RestoreHpStep) < h.MaxHP {
		h.HealthPoint = h.MaxHP
	}

	h.HealthPoint += h.RestoreHpStep
}

type AI[HP constraints.Integer] struct {
	HealthPoint HP
	MaxHP       HP
}

func (a *AI[HP]) getHP() HP {
	return a.getHP()
}

func (a *AI[HP]) restoreHP() {
	a.HealthPoint = a.MaxHP
}

func createFighter[HP constraints.Integer](fighterType FighterType, maxHp, restoreStep HP) (Fighter, error) {
	switch fighterType {
	case ai:
		return
	}
}
