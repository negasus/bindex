package bindex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	bi := New()

	bi.Set("KEY", []int{1, 65})

	v, ok := bi.v["KEY"]
	assert.True(t, ok)
	assert.Len(t, v, 2)
	assert.Equal(t, uint64(1), v[0])
	assert.Equal(t, uint64(1), v[1])
}

func TestAnd(t *testing.T) {
	bi := New()

	bi.Set("KEY_1", []int{1, 2, 3, 4, 5, 10, 20, 30, 40})
	bi.Set("KEY_2", []int{2, 3, 40, 50, 100})
	bi.Set("KEY_3", []int{1, 3, 50, 20})

	result1 := bi.Select("KEY_1").Result()
	assert.Len(t, result1, 9)
	assert.Equal(t, 1, result1[0])
	assert.Equal(t, 2, result1[1])
	assert.Equal(t, 3, result1[2])
	assert.Equal(t, 4, result1[3])
	assert.Equal(t, 5, result1[4])
	assert.Equal(t, 10, result1[5])
	assert.Equal(t, 20, result1[6])
	assert.Equal(t, 30, result1[7])
	assert.Equal(t, 40, result1[8])

	result2 := bi.Select("KEY_1").And("KEY_2").Result()
	assert.Len(t, result2, 3)
	assert.Equal(t, 2, result2[0])
	assert.Equal(t, 3, result2[1])
	assert.Equal(t, 40, result2[2])

	result3 := bi.Select("KEY_1").And("KEY_2").And("KEY_3").Result()
	assert.Len(t, result3, 1)
	assert.Equal(t, 3, result3[0])

	result4 := bi.Select("NO_KEY").Result()
	assert.Len(t, result4, 0)

	result5 := bi.Select("KEY_1").And("NO_KEY").Result()
	assert.Len(t, result5, 0)

}

func TestAndNot(t *testing.T) {
	bi := New()

	bi.Set("KEY_1", []int{1, 2, 3, 4, 5, 10, 20, 30, 40})
	bi.Set("KEY_2", []int{2, 3, 40, 50, 100})

	result1 := bi.Select("KEY_1").AndNot("KEY_2").Result()
	assert.Len(t, result1, 6)
	assert.Equal(t, 1, result1[0])
	assert.Equal(t, 4, result1[1])
	assert.Equal(t, 5, result1[2])
	assert.Equal(t, 10, result1[3])
	assert.Equal(t, 20, result1[4])
	assert.Equal(t, 30, result1[5])

	result2 := bi.Select("KEY_2").AndNot("NO_KEY").Result()
	assert.Len(t, result2, 5)
	assert.Equal(t, 2, result2[0])
	assert.Equal(t, 3, result2[1])
	assert.Equal(t, 40, result2[2])
	assert.Equal(t, 50, result2[3])
	assert.Equal(t, 100, result2[4])

	result3 := bi.Select("NO_KEY").AndNot("NO_KEY").Result()
	assert.Len(t, result3, 0)
}

func TestOr(t *testing.T) {
	bi := New()

	bi.Set("KEY_1", []int{1, 3, 10, 20, 30, 40})
	bi.Set("KEY_2", []int{2, 3, 40, 50, 100})

	result1 := bi.Select("KEY_1").Or("KEY_2").Result()
	assert.Len(t, result1, 9)
	assert.Equal(t, 1, result1[0])
	assert.Equal(t, 2, result1[1])
	assert.Equal(t, 3, result1[2])
	assert.Equal(t, 10, result1[3])
	assert.Equal(t, 20, result1[4])
	assert.Equal(t, 30, result1[5])
	assert.Equal(t, 40, result1[6])
	assert.Equal(t, 50, result1[7])
	assert.Equal(t, 100, result1[8])

	result2 := bi.Select("KEY_2").AndNot("NO_KEY").Result()
	assert.Len(t, result2, 5)
	assert.Equal(t, 2, result2[0])
	assert.Equal(t, 3, result2[1])
	assert.Equal(t, 40, result2[2])
	assert.Equal(t, 50, result2[3])
	assert.Equal(t, 100, result2[4])

	result3 := bi.Select("NO_KEY").AndNot("NO_KEY").Result()
	assert.Len(t, result3, 0)
}

func TestComplex(t *testing.T) {
	bi := New()

	bi.Set("KEY_1", []int{1, 3, 10, 20, 30, 40})
	bi.Set("KEY_2", []int{2, 3, 10, 40, 50, 100})
	bi.Set("KEY_3", []int{2, 3, 50, 100})
	bi.Set("KEY_4", []int{2, 40})

	result := bi.Select("KEY_1").And("KEY_2").AndNot("KEY_3").Or("KEY_4").Result()
	assert.Len(t, result, 3)
	assert.Equal(t, 2, result[0])
	assert.Equal(t, 10, result[1])
	assert.Equal(t, 40, result[2])
}
