package figther

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type fighterCreateData[HP HeathPoint] struct {
	fighterType FighterType
	maxHp       HP
	restoreStep HP
}

type expect struct {
	result Fighter[uint]
	error  string
}

func TestCreateFighter(t *testing.T) {
	type test struct {
		name              string
		fighterCreateData fighterCreateData[uint8]
		expect            expect
	}

	tests := []test{
		{
			name: `create human`,
			fighterCreateData: fighterCreateData[uint8]{
				fighterType: `human`,
				maxHp:       100,
				restoreStep: 20,
			},
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
			name: `create ai`,
			fighterCreateData: fighterCreateData[uint8]{
				fighterType: `ai`,
				maxHp:       100,
				restoreStep: 0,
			},
			expect: expect{
				error: "",
				result: &AI[uint]{
					Hp:    100,
					MaxHp: 100,
				},
			},
		},
		{
			name: `wrong type`,
			fighterCreateData: fighterCreateData[uint8]{
				fighterType: `cyborg`,
				maxHp:       100,
				restoreStep: 20,
			},
			expect: expect{
				error: "wrong type",
			},
		},
	}

	testFunc := func(tc test) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			f, err := createFighter(tc.fighterCreateData.fighterType, tc.fighterCreateData.maxHp, tc.fighterCreateData.restoreStep)
			assert.Equal(t, tc.expect.result, f)
			if tc.expect.error == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expect.error)
			}
		}
	}
	for _, tc := range tests {
		t.Run(tc.name, testFunc(tc))
	}
}

type testGetHp[HP HeathPoint] struct {
	name              string
	fighterCreateData fighterCreateData[HP]
	expect            HP
}

func TestGetHp(t *testing.T) {
	testsUint8 := []testGetHp[uint8]{
		{
			name: `human uint8 getHP`,
			fighterCreateData: fighterCreateData[uint8]{
				fighterType: `human`,
				maxHp:       111,
				restoreStep: 20,
			},
			expect: 111,
		},
		{
			name: `ai uint8 getHP`,
			fighterCreateData: fighterCreateData[uint8]{
				fighterType: `ai`,
				maxHp:       123,
				restoreStep: 20,
			},
			expect: 123,
		},
	}
	testsUint16 := []testGetHp[uint16]{
		{
			name: `human uint16 getHP`,
			fighterCreateData: fighterCreateData[uint16]{
				fighterType: `human`,
				maxHp:       111,
				restoreStep: 20,
			},
			expect: 111,
		},
		{
			name: `ai uint16 getHP`,
			fighterCreateData: fighterCreateData[uint16]{
				fighterType: `ai`,
				maxHp:       123,
				restoreStep: 20,
			},
			expect: 123,
		},
	}

	for _, tc := range testsUint8 {
		t.Run(tc.name, testGetHP(tc))
	}

	for _, tc := range testsUint16 {
		t.Run(tc.name, testGetHP(tc))
	}
}

func testGetHP[HP HeathPoint](tc testGetHp[HP]) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		f, err := createFighter(tc.fighterCreateData.fighterType, tc.fighterCreateData.maxHp, tc.fighterCreateData.restoreStep)
		assert.NoError(t, err)
		assert.Equal(t, tc.expect, f.getHp())
	}
}
