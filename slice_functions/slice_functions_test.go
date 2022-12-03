package slice_functions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	l := GroupBy(a, func(s string) (int, string) { return len(s), s })
	assert.Equal(t, New("a", "b", "l"), l[1])
	assert.Equal(t, New("de"), l[2])
	assert.Equal(t, New("abc", "hjk"), l[3])
	assert.Equal(t, 0, len(l[4]))
}

func TestMapBy(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	mapped := MapBy(a, func(s string) (string, string) { return s, s })
	assert.Equal(t, "l", mapped["l"])
	assert.Equal(t, "de", mapped["de"])
	assert.Equal(t, "hjk", mapped["hjk"])
	assert.Equal(t, "", mapped["notfound"])
}

func TestMap(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	l := Map(a, func(s string) int { return len(s) })
	assert.Equal(t, New(1, 3, 2, 3, 1, 1), l)
}

func TestFlatMap(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	l := FlatMap(a, func(s string) []int { return New(len(s), 0) })
	assert.Equal(t, New(1, 0, 3, 0, 2, 0, 3, 0, 1, 0, 1, 0), l)
}

func TestFlatTry(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	l, err := FlatMapTry(a, func(s string) ([]int, error) { return New(len(s), 0), nil })
	assert.NoError(t, err)
	assert.Equal(t, New(1, 0, 3, 0, 2, 0, 3, 0, 1, 0, 1, 0), l)

	_, err = FlatMapTry(a, func(s string) ([]int, error) { return New(len(s), 0), errors.New("some error") })
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
}

func TestFilter(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	oddLength := Filter(a, func(s string) bool { return len(s)%2 == 1 })
	assert.Equal(t, New("a", "abc", "hjk", "b", "l"), oddLength)
}

func TestForEach(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	length := 0
	ForEach(a, func(s string) { length++ })
	assert.Equal(t, 6, length)
}

func TestForAll(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	assert.True(t, ForAll(a, func(s string) bool { return len(s) < 10 }))
	assert.True(t, ForAll([]string{}, func(s string) bool { return len(s) < 10 }))
	assert.False(t, ForAll(a, func(s string) bool { return len(s) < 3 }))
}

func TestExists(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}

	assert.False(t, Exists(a, func(s string) bool { return len(s) > 10 }))
	assert.False(t, Exists([]string{}, func(s string) bool { return len(s) > 10 }))
	assert.True(t, Exists(a, func(s string) bool { return len(s) < 3 }))
}

func TestCount(t *testing.T) {
	a := []string{"a", "abc", "de", "hjk", "b", "l"}
	assert.Equal(t, 2, Count(a, func(s string) bool { return len(s) == 3 }))
	assert.Equal(t, 1, Count(a, func(s string) bool { return len(s) == 2 }))
	assert.Equal(t, 0, Count(a, func(s string) bool { return len(s) == 5 }))
}
