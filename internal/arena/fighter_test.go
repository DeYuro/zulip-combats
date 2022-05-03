package arena

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFighter[HP HeathPoint](t *testing.T) {
	type expect[HP HeathPoint] struct {
		result Fighter[HP]
		error  string
	}

	type test struct {
		FighterType FighterType
		maxHp       uint
		restoreStep uint
		expect      expect
	}

	testCases := []test{
		{
			`human`,
			100,
			20,
			expect{
				error: "",
				result: Human[HP]{
					100,
					100,
					20,
				},
			},
		},
	}

	t.Run("create", func(t *testing.T) {
		for _, v := range testCases {
			f, err := createFighter(v.FighterType, v.maxHp, v.restoreStep)
			assert.Equal(t, v.expect.result, f)
			assert.NoError(t, err)
		}
	})
}
