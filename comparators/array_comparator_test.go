package comparators

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
)

func TestStringArraysAreEqualIfBothAreNil(t *testing.T) {
	assert.True(t, CompareStringArrays(nil, nil))
}

func TestStringArraysAreNotEqualIfOnlyOneIsNil(t *testing.T) {
	assert.False(t, CompareStringArrays([]string{}, nil))
}

func TestStringArraysAreNotEqualIfTheirLengthIsDifferent(t *testing.T) {
	assert.False(t, CompareStringArrays([]string{}, []string{"one-element"}))
}

func TestStringArraysAreNotEqualIfTheirElementsDifferent(t *testing.T) {
	assert.False(t, CompareStringArrays([]string{"other-element"}, []string{"one-element"}))
}

func TestStringArraysAreEqualIfAllElementsAreEqual(t *testing.T) {
	assert.True(t, CompareStringArrays([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
}