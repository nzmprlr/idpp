package idpp

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewID8(t *testing.T) {
	t.Parallel()

	id1 := NewID8()
	id2 := NewID8()

	assert.NoError(t, id1.Validate())
	assert.NoError(t, id2.Validate())

	pid1, err := ParseID8(id1.String())
	assert.NoError(t, err)
	assert.Equal(t, id1.String(), pid1.String())

	pid2, err := ParseID8Hex(id2.Hex())
	assert.NoError(t, err)
	assert.Equal(t, id2.Hex(), pid2.Hex())

	uid1, err := strconv.ParseUint(id1.String(), 10, 64)
	assert.NoError(t, err)

	uid2, err := strconv.ParseUint(id2.String(), 10, 64)
	assert.NoError(t, err)

	assert.True(t, uid1 < uid2)
	assert.True(t, id1.Time().Unix() <= id2.Time().Unix())
	assert.True(t, id1.Time().UnixNano() <= id2.Time().UnixNano())

	assert.Equal(t, id1.WorkerID(), id2.WorkerID())
	assert.True(t, id1.Sequence() < id2.Sequence() || id1.Sequence() == math.MaxUint16)
}

func TestNewID8WithTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		t     time.Time
		valid bool
	}{
		{time.Now(), true},
		{time.Unix(0, int64(timeEpoch)), false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.t.String(), func(t *testing.T) {
			t.Parallel()

			id := NewID8WithTime(tc.t)

			if tc.valid {
				assert.NoError(t, id.Validate())
				assert.Equal(t, tc.t.Unix(), id.Time().Unix())
			} else {
				assert.Error(t, id.Validate())
			}
		})
	}
}

func TestParseID8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s     string
		valid bool
	}{
		{NewID8().String(), true},
		{"", false},
		{"not a number", false},
		{strings.Repeat("#", 17), false},
		{strings.Repeat("1", 16), true},
		{"1", false},
		{"13245", false},
		{"1cox", false},
		{"fffe", false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.s, func(t *testing.T) {
			t.Parallel()

			id, err := ParseID8(tc.s)

			if tc.valid {
				assert.NoError(t, err)
				assert.NoError(t, id.Validate())
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestParseID8Hex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s     string
		valid bool
	}{
		{NewID8().Hex(), true},
		{"", false},
		{"not a number", false},
		{strings.Repeat("#", 17), false},
		{strings.Repeat("a", 16), true},
		{"1", false},
		{"13245", false},
		{"1cox", false},
		{"fffe", false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.s, func(t *testing.T) {
			t.Parallel()

			id, err := ParseID8Hex(tc.s)

			if tc.valid {
				assert.NoError(t, err)
				assert.NoError(t, id.Validate())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
