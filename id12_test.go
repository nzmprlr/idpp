package idpp

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewID12(t *testing.T) {
	t.Parallel()

	id1 := NewID12()
	id2 := NewID12()

	assert.NoError(t, id1.Validate())
	assert.NoError(t, id2.Validate())

	pid1, err := ParseID12(id1.String())
	assert.NoError(t, err)
	assert.Equal(t, id1.String(), pid1.String())

	pid2, err := ParseID12Hex(id2.Hex())
	assert.NoError(t, err)
	assert.Equal(t, id2.Hex(), pid2.Hex())

	assert.True(t, id1.Time().Unix() <= id2.Time().Unix())
	assert.True(t, id1.Time().UnixNano() <= id2.Time().UnixNano())

	assert.Equal(t, id1.WorkerID(), id2.WorkerID())
	assert.True(t, id1.Sequence() < id2.Sequence() || id1.Sequence() == math.MaxUint16)
}

func TestNewID12WithTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		t     time.Time
		valid bool
	}{
		{time.Now(), true},
		{time.Unix(0, int64(timeEpoch)), true},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.t.String(), func(t *testing.T) {
			t.Parallel()

			id := NewID12WithTime(tc.t)

			if tc.valid {
				assert.NoError(t, id.Validate())
				assert.Equal(t, tc.t.Unix(), id.Time().Unix())
			} else {
				assert.Error(t, id.Validate())
			}
		})
	}
}

func TestParseID12(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s     string
		valid bool
	}{
		{NewID12().String(), true},
		{NewID12().Hex(), true},
		{"", false},
		{"not a number", false},
		{strings.Repeat("#", 25), false},
		{strings.Repeat("a", 24), true},
		{"1", true},
		{"13245", true},
		{"1cox", false},
		{"fffe", true},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.s, func(t *testing.T) {
			t.Parallel()

			id, err := ParseID12(tc.s)

			if tc.valid {
				assert.NoError(t, err)
				assert.NoError(t, id.Validate())
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestParseID12Hex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s     string
		valid bool
	}{
		{NewID12().String(), true},
		{NewID12().Hex(), true},
		{"", false},
		{"not a number", false},
		{strings.Repeat("#", 25), false},
		{strings.Repeat("a", 24), true},
		{"1", true},
		{"13245", true},
		{"1cox", false},
		{"fffe", true},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.s, func(t *testing.T) {
			t.Parallel()

			id, err := ParseID12Hex(tc.s)

			if tc.valid {
				assert.NoError(t, err)
				assert.NoError(t, id.Validate())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
