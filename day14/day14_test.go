package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamples(t *testing.T) {
	out := doElves(9)
	assert.Equal(t, "5158916779", out)

	out = doElves(5)
	assert.Equal(t, "0124515891", out)

	out = doElves(18)
	assert.Equal(t, "9251071085", out)

	out = doElves(2018)
	assert.Equal(t, "5941429882", out)
}
