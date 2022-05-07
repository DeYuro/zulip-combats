package figther

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CreateFighterTest struct {
	name        string
	fighterType FighterType
	maxHp       uint
	restoreStep uint
	expect      expect
}

type expect struct {
	result Fighter[uint]
	error  string
}

func getTestcases() []CreateFighterTest {
	return []CreateFighterTest{
		{
			name:        `create human`,
			fighterType: `human`,
			maxHp:       100,
			restoreStep: 20,
			expect: expect{
				error: "",
				result: &Human[uint]{
					Hp:          100,
					MaxHp:       100,
					RestoreStep: 20,
				},
			},
		},
		{
			name:        `create ai`,
			fighterType: `ai`,
			maxHp:       100,
			restoreStep: 0,
			expect: expect{
				error: "",
				result: &AI[uint]{
					Hp:    100,
					MaxHp: 100,
				},
			},
		},
		{
			name:        `wrong type`,
			fighterType: `cyborg`,
			maxHp:       100,
			restoreStep: 20,
			expect: expect{
				error: "wrong type",
			},
		},
	}
}

func TestCreateFighter(t *testing.T) {
	tests := getTestcases()

	for _, tc := range tests {
		t.Run(tc.name, runCreateFighterTestCases(tc))
	}
}

func runCreateFighterTestCases(tc CreateFighterTest) func(t *testing.T) {
	return func(t *testing.T) {
		f, err := createFighter(tc.fighterType, tc.maxHp, tc.restoreStep)
		assert.Equal(t, tc.expect.result, f)
		if tc.expect.error == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, tc.expect.error)
		}
	}
}

func TestGetHp(t *testing.T) {

}
