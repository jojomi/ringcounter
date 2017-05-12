package ringcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingcounterBasic(t *testing.T) {
	r := New(10, 5)
	assert.Equal(t, 5, r.Value, "Initial value should be 5")

	r.Add(2)
	assert.Equal(t, 7, r.Value, "Value should be incremented")

	r.Advance()
	assert.Equal(t, 8, r.Value, "Value should be incremented")

	r.Add(-4)
	assert.Equal(t, 4, r.Value, "Value should be incremented")
}

func TestRingcounterOverflow(t *testing.T) {
	r := New(7, 0)

	r.Add(48)
	assert.Equal(t, 6, r.Value, "Value should have seen an overflow")

	r.Advance()
	assert.Equal(t, 0, r.Value, "Value should have seen an overflow")

	r.Add(-48)
	assert.Equal(t, 1, r.Value, "Value should have seen an underflow")
}
