package idpp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

// id12 describes a 12byte(96bit) time based id
type id12 [12]byte

func (id id12) String() string { return id.Hex() }

// Hex returns the ID in hex format.
func (id id12) Hex() string { return strings.TrimLeft(hex.EncodeToString(id[:]), "0") }

// Time returns the timestamp encoded into the ID.
func (id id12) Time() time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint64(id[:8])+timeEpoch))
}

// WorkerID returns the worker ID associated with the ID.
func (id id12) WorkerID() uint16 { return binary.BigEndian.Uint16(id[8:10]) }

// Sequence returns the sequence number associated with the ID.
func (id id12) Sequence() uint16 { return binary.BigEndian.Uint16(id[10:]) }

// Validate returns error if the ID is invalid.
func (id id12) Validate() error { return nil } // seems no case to validate.

// NewID12 creates a new 12byte ID from the current UTC time.
func NewID12() ID { return NewID12WithTime(time.Now().UTC()) }

// NewID12WithTime creates a new 12byte ID with the specified time.
func NewID12WithTime(t time.Time) ID {
	var id id12

	binary.BigEndian.PutUint64(id[:8], uint64(t.UnixNano())-timeEpoch)
	binary.BigEndian.PutUint16(id[8:10], uint16(workerID))
	binary.BigEndian.PutUint16(id[10:], uint16(atomic.AddUint32(&sequence, 1)))

	return id
}

// ParseID12 parses a 12byte ID from a string.
func ParseID12(s string) (ID, error) { return ParseID12Hex(s) }

// ParseID12Hex parses a 12byte ID from a hex string.
func ParseID12Hex(h string) (ID, error) {
	var id id12

	l := len(h)
	if l == 0 || l > 24 {
		return id, fmt.Errorf("invalid ID: %s", h)
	} else if l&1 == 1 { // check hex is odd, if odd add zero to head
		h = "0" + h
	}

	d, err := hex.DecodeString(h)
	if err != nil {
		return id, err
	}

	// get bytes from decoded(d) id
	l = len(d)
	for i := 0; i < l; i++ {
		id[i] = d[i]
	}

	return id, id.Validate()
}
