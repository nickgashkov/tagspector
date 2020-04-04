package slices

import (
	"github.com/nickgashkov/tagspector/testing/assert"
	"testing"
)

func TestTwoEqualSlices(t *testing.T) {
	// Given
	sliceA := []string{"1", "2", "3"}
	sliceB := []string{"1", "2", "3"}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, isEqual, "Expected slices to be equal")
}

func TestTwoEqualNilSlices(t *testing.T) {
	// Given
	var sliceA []string
	var sliceB []string

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, isEqual, "Expected nil slices to be equal")
}

func TestTwoEqualEmptySlices(t *testing.T) {
	// Given
	sliceA := []string{}
	sliceB := []string{}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, isEqual, "Expected empty slices to be equal")
}

func TestNilAndEmptySlicesAreEqual(t *testing.T) {
	// Given
	var sliceA []string
	var sliceB = []string{}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, isEqual, "Expected nil and empty slices to be equal")
}

func TestTwoNotEqualSlicesOfTheSameLength(t *testing.T) {
	// Given
	sliceA := []string{"1", "2", "3"}
	sliceB := []string{"4", "5", "6"}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, !isEqual, "Expected slices not to be equal")
}

func TestTwoNotEqualSlicesOfDifferentLength(t *testing.T) {
	// Given
	sliceA := []string{"1", "2"}
	sliceB := []string{"1"}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, !isEqual, "Expected uneven slices not to be equal")
}

func TestSliceAndNilSliceAreNotEqual(t *testing.T) {
	// Given
	var sliceA = []string{"1"}
	var sliceB []string

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, !isEqual, "Expected regular and nil slices not to be equal")
}

func TestSliceAndEmptySliceAreNotEqual(t *testing.T) {
	// Given
	sliceA := []string{"1"}
	sliceB := []string{}

	// When
	isEqual := StringsEqual(sliceA, sliceB)
	isEqualReversed := StringsEqual(sliceB, sliceA)

	// Then
	assert.Equal(t, isEqual, isEqualReversed)
	assert.True(t, !isEqual, "Expected regular and empty slices not to be equal")
}
