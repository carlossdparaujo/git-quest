package comparators

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
)

func TestStringArraysAreEqualIfBothAreNil(t *testing.T) {
	assert.True(t, compareStringArrays(nil, nil))
}

func TestStringArraysAreNotEqualIfOnlyOneIsNil(t *testing.T) {
	assert.False(t, compareStringArrays([]string{}, nil))
}

func TestStringArraysAreNotEqualIfTheirLengthIsDifferent(t *testing.T) {
	assert.False(t, compareStringArrays([]string{}, []string{"one-element"}))
}

func TestStringArraysAreNotEqualIfTheirElementsDifferent(t *testing.T) {
	assert.False(t, compareStringArrays([]string{"other-element"}, []string{"one-element"}))
}

func TestStringArraysAreEqualIfAllElementsAreEqual(t *testing.T) {
	assert.True(t, compareStringArrays([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
}