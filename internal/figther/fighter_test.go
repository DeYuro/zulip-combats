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
	result Fighter[int]
	error  string
}

func TestCreateFighter(t *testing.T) {
	type test struct {
		name              string
		fighterCreateData fighterCreateData[int8]
		expect            expect
	}

	tests := []test{
		{
			name: `create human`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `human`,
				maxHp:       100,
				restoreStep: 20,
			},
			expect: expect{
				error: "",
				result: &Human[int]{
					Hp:          100,
					MaxHp:       100,
					RestoreStep: 20,
				},
			},
		},
		{
			name: `create ai`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `ai`,
				maxHp:       100,
				restoreStep: 0,
			},
			expect: expect{
				error: "",
				result: &AI[int]{
					Hp:    100,
					MaxHp: 100,
				},
			},
		},
		{
			name: `wrong type`,
			fighterCreateData: fighterCreateData[int8]{
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

type testGetHP[HP HeathPoint] struct {
	name              string
	fighterCreateData fighterCreateData[HP]
	expect            HP
}

func TestGetHp(t *testing.T) {
	testsUint8 := []testGetHP[int8]{
		{
			name: `human int8 getHP`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `human`,
				maxHp:       111,
				restoreStep: 20,
			},
			expect: 111,
		},
		{
			name: `ai int8 getHP`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `ai`,
				maxHp:       123,
				restoreStep: 20,
			},
			expect: 123,
		},
	}
	testsUint16 := []testGetHP[int16]{
		{
			name: `human int16 getHP`,
			fighterCreateData: fighterCreateData[int16]{
				fighterType: `human`,
				maxHp:       111,
				restoreStep: 20,
			},
			expect: 111,
		},
		{
			name: `ai int16 getHP`,
			fighterCreateData: fighterCreateData[int16]{
				fighterType: `ai`,
				maxHp:       123,
				restoreStep: 20,
			},
			expect: 123,
		},
	}

	for _, tc := range testsUint8 {
		t.Run(tc.name, runTestGetHP(tc))
	}

	for _, tc := range testsUint16 {
		t.Run(tc.name, runTestGetHP(tc))
	}
}

func runTestGetHP[HP HeathPoint](tc testGetHP[HP]) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		f, err := createFighter(tc.fighterCreateData.fighterType, tc.fighterCreateData.maxHp, tc.fighterCreateData.restoreStep)
		assert.NoError(t, err)
		assert.Equal(t, tc.expect, f.getHp())
	}
}

type testMethodsHP[HP HeathPoint] struct {
	name              string
	fighterCreateData fighterCreateData[HP]
	damage            HP
	expect            HP
}

func TestTakeDamage(t *testing.T) {
	testsUint := []testMethodsHP[int]{
		{
			name: `human int takeDamage`,
			fighterCreateData: fighterCreateData[int]{
				fighterType: `human`,
				maxHp:       100,
				restoreStep: 20,
			},
			damage: 10,
			expect: 90,
		},
		{
			name: `ai int takeDamage`,
			fighterCreateData: fighterCreateData[int]{
				fighterType: `ai`,
				maxHp:       12,
				restoreStep: 12,
			},
			damage: 20,
			expect: 0,
		},
	}
	testsUint32 := []testMethodsHP[int32]{
		{
			name: `human int32 takeDamage`,
			fighterCreateData: fighterCreateData[int32]{
				fighterType: `human`,
				maxHp:       30,
				restoreStep: 20,
			},
			damage: 40,
			expect: 0,
		},
		{
			name: `ai int32 takeDamage`,
			fighterCreateData: fighterCreateData[int32]{
				fighterType: `ai`,
				maxHp:       123,
				restoreStep: 20,
			},
			damage: 23,
			expect: 100,
		},
	}

	for _, tc := range testsUint {
		t.Run(tc.name, runTestTakeDamage(tc))
	}

	for _, tc := range testsUint32 {
		t.Run(tc.name, runTestTakeDamage(tc))
	}
}

func runTestTakeDamage[HP HeathPoint](tc testMethodsHP[HP]) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		f, err := createFighter(tc.fighterCreateData.fighterType, tc.fighterCreateData.maxHp, tc.fighterCreateData.restoreStep)
		assert.NoError(t, err)
		f.takeDamage(tc.damage)
		assert.Equal(t, tc.expect, f.getHp())
	}
}

func TestRestoreHP(t *testing.T) {
	testsUint64 := []testMethodsHP[int64]{
		{
			name: `human int64 restoreHP`,
			fighterCreateData: fighterCreateData[int64]{
				fighterType: `human`,
				maxHp:       120,
				restoreStep: 10,
			},
			damage: 30,
			expect: 100,
		},
		{
			name: `ai int64 restoreHP`,
			fighterCreateData: fighterCreateData[int64]{
				fighterType: `ai`,
				maxHp:       120,
				restoreStep: 12,
			},
			damage: 100,
			expect: 120,
		},
	}
	testsUint8 := []testMethodsHP[int8]{
		{
			name: `human int8 takeDamage`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `human`,
				maxHp:       30,
				restoreStep: 20,
			},
			damage: 1,
			expect: 30,
		},
		{
			name: `ai int8 takeDamage`,
			fighterCreateData: fighterCreateData[int8]{
				fighterType: `ai`,
				maxHp:       100,
				restoreStep: 0,
			},
			damage: 50,
			expect: 100,
		},
	}

	for _, tc := range testsUint64 {
		t.Run(tc.name, runTestRestoreHP(tc))
	}

	for _, tc := range testsUint8 {
		t.Run(tc.name, runTestRestoreHP(tc))
	}
}

func runTestRestoreHP[HP HeathPoint](tc testMethodsHP[HP]) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		f, err := createFighter(tc.fighterCreateData.fighterType, tc.fighterCreateData.maxHp, tc.fighterCreateData.restoreStep)
		assert.NoError(t, err)
		f.takeDamage(tc.damage)
		f.restoreHp()
		assert.Equal(t, tc.expect, f.getHp())
	}
}
